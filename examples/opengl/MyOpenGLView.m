#import "MyOpenGLView.h"

#import <OpenGL/OpenGL.h>

void deleteGLRenderFunc(void * obj);
bool callGLRenderFunc(void * obj);

@implementation MyOpenGLView

- (void)awakeFromNib
{
    NSOpenGLPixelFormatAttribute pixFmtAttrs[] = {
		NSOpenGLPFAOpenGLProfile, NSOpenGLProfileVersion4_1Core,
		NSOpenGLPFAColorSize, 24,
		NSOpenGLPFAAlphaSize, 8,
		NSOpenGLPFADoubleBuffer,
		NSOpenGLPFAAccelerated,
		0,
    };

    NSOpenGLPixelFormat *pixFmt = [[NSOpenGLPixelFormat alloc] initWithAttributes:pixFmtAttrs];
    super.pixelFormat = pixFmt;
    [pixFmt release];
}

- (void)drawRect:(NSRect)r
{
    [super drawRect:r];

    if (callGLRenderFunc(self)) {
        [self.openGLContext flushBuffer];
    }
}

- (void)dealloc
{
    deleteGLRenderFunc(self);
    [super dealloc];
}

@end