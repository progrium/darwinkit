#!/bin/bash

# analyze_technologies.sh - Analyzes Apple's technologies.json file
# This script downloads and analyzes Apple's technologies.json file to help
# with framework implementation planning.

OUTPUT_DIR="./tech_analysis"
ANALYSIS_FILE="$OUTPUT_DIR/framework_analysis.md"
COVERAGE_FILE="$OUTPUT_DIR/coverage_report.md"

# Create output directory
mkdir -p "$OUTPUT_DIR"

echo "Analyzing Apple frameworks and technologies..."

# Run the tech_analyzer tool
go run cmd/tools/tech_analyzer/main.go \
  --outdir="$OUTPUT_DIR" \
  --count-symbols \
  --list-macos \
  --output="$ANALYSIS_FILE"

echo "Generating implementation coverage report..."

# Get list of implemented frameworks from macos directory
IMPLEMENTED_FRAMEWORKS=$(find ./macos -maxdepth 1 -type d | grep -v "_examples\|_wip" | sed 's/\.\/macos\///' | tr '\n' ',' | sed 's/,$//')

# Run the tool again to get specifically implemented frameworks
go run cmd/tools/tech_analyzer/main.go \
  --outdir="$OUTPUT_DIR" \
  --frameworks="$IMPLEMENTED_FRAMEWORKS" \
  --count-symbols > /dev/null

# Create a coverage report
cat <<EOF > "$COVERAGE_FILE"
# DarwinKit Framework Implementation Coverage

This report shows the current implementation status of Apple frameworks in DarwinKit.

## Overview

DarwinKit currently implements frameworks from the following technology areas:
- Core frameworks (Foundation, CoreFoundation)
- UI frameworks (AppKit)
- Graphics (CoreGraphics, CoreImage)
- Media (AVFoundation)
- Data (CoreData)
- Networking and Cloud (CloudKit)
- Location and Maps (CoreLocation, MapKit)
- Health and Fitness (HealthKit)
- Home Automation (HomeKit)
- Gaming (GameKit)
- Security (Security)
- Scripting (JavaScriptCore)
- Notifications (UserNotifications)

## Implementation Status by Category

| Category | Implemented | Total Available | Coverage % |
|----------|-------------|-----------------|------------|
EOF

# Parse the technologies.json to get category stats
python3 -c "
import json
import os
import re

# Read technologies.json
with open('$OUTPUT_DIR/technologies.json', 'r') as f:
    data = json.load(f)

# Get implemented frameworks
implemented = '${IMPLEMENTED_FRAMEWORKS}'.split(',')

# Create category mapping
categories = {}
framework_to_category = {}

for tech in data['technologies']:
    if tech['type'] == 'category':
        categories[tech['id']] = {
            'name': tech['title'],
            'implemented': 0,
            'total': 0
        }
    elif tech['type'] == 'framework':
        # Find parent category
        for parent_tech in data['technologies']:
            if parent_tech.get('children') and tech['id'] in parent_tech.get('children', []):
                category_id = parent_tech['id']
                framework_to_category[tech['id']] = category_id
                categories.setdefault(category_id, {'name': parent_tech['title'], 'implemented': 0, 'total': 0})
                categories[category_id]['total'] += 1
                
                # Check if this framework is implemented
                for impl in implemented:
                    if re.match(f'^{impl}$', tech['title'], re.IGNORECASE):
                        categories[category_id]['implemented'] += 1
                        break
                break

# Output category stats
for category_id, stats in categories.items():
    if stats['total'] > 0:
        coverage = (stats['implemented'] / stats['total']) * 100
        print(f\"| {stats['name']} | {stats['implemented']} | {stats['total']} | {coverage:.1f}% |\")
" >> "$COVERAGE_FILE"

# Add detailed framework section
cat <<EOF >> "$COVERAGE_FILE"

## Recently Added Frameworks

The following frameworks have been recently added to DarwinKit:

- EventKit: Calendar and reminder management
- MapKit: Map and location services
- HealthKit: Health data access and monitoring
- HomeKit: Home automation control
- GameKit: Game Center integration
- Security: Keychain and cryptographic operations
- JavaScriptCore: JavaScript execution
- UserNotifications: System notification management

## Next Framework Recommendations

Based on symbol count and platform importance, the following frameworks are recommended for implementation:

EOF

# Add recommendations based on symbol count
python3 -c "
import json

# Read technologies.json
with open('$OUTPUT_DIR/technologies.json', 'r') as f:
    data = json.load(f)

# Get implemented frameworks
implemented = '${IMPLEMENTED_FRAMEWORKS}'.split(',')

# Find frameworks not yet implemented with high symbol counts
candidates = []
for tech in data['technologies']:
    if tech['type'] == 'framework' and tech['symbolCount'] > 100:
        is_implemented = False
        for impl in implemented:
            if impl.lower() == tech['title'].lower():
                is_implemented = True
                break
        
        if not is_implemented and 'macos' in [p.lower() for p in tech.get('platforms', [])]:
            candidates.append({
                'title': tech['title'],
                'symbols': tech['symbolCount'],
                'platforms': tech.get('platforms', [])
            })

# Sort by symbol count
candidates.sort(key=lambda x: x['symbols'], reverse=True)

# Output top 10 recommendations
for i, candidate in enumerate(candidates[:10]):
    platforms = ', '.join(candidate['platforms'])
    print(f\"{i+1}. **{candidate['title']}** - {candidate['symbols']} symbols ({platforms})\")
" >> "$COVERAGE_FILE"

echo "Analysis complete!"
echo "Framework analysis report: $ANALYSIS_FILE"
echo "Coverage report: $COVERAGE_FILE"