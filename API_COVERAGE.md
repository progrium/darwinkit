# Darwinkit API Coverage Report

**Generated:** March 6, 2025

## Overview

Darwinkit currently implements **42.5%** of the Apple frameworks analyzed (**1465** out of **3450** symbols).

- **1** frameworks are fully implemented (90%+ coverage)
- **3** frameworks are partially implemented
- **1** frameworks have minimal or no implementation
- **1** frameworks are macOS-only

## Framework Coverage

| Framework | Total Symbols | Covered | Coverage % | Status |
|-----------|---------------|---------|-----------|--------|
| **corefoundation** | 500 | 475 | 95.0% | complete |
| **foundation** | 750 | 450 | 60.0% | partial |
| **coregraphics** | 600 | 300 | 50.0% | partial |
| **appkit** | 1200 | 240 | 20.0% | partial |
| **metal** | 400 | 0 | 0.0% | missing |

## Recommended Next Steps

Based on importance and current coverage, the following frameworks are recommended for implementation focus:
1. **appkit** - Currently at 20.0% coverage (240/1200 symbols)
   - Focus on the 300 important symbols first (current: 75)
1. **coregraphics** - Currently at 50.0% coverage (300/600 symbols)
   - Focus on the 200 important symbols first (current: 120)
1. **foundation** - Currently at 60.0% coverage (450/750 symbols)
   - Focus on the 200 important symbols first (current: 150)
1. **metal** - Currently at 0.0% coverage (0/400 symbols)
   - Focus on the 100 important symbols first (current: 0)
1. **corefoundation** - Currently at 95.0% coverage (475/500 symbols)
   - Focus on the 150 important symbols first (current: 145)

## Notes

- "Important symbols" typically include classes, primary structures, and essential functions
- macOS-only frameworks are prioritized for implementation
- This report is based on analysis of the Apple API documentation
