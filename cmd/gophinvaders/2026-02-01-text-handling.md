In order to display text, we need some config governing text size. We use `golang.org/x/image/font/opentype` fonts. For simplicity, we assume 72 DPI, but we may refine this later, so create config for this.

```go
pointsPerInch = 72
DPI = 72
pointPerPixel float32 = pointsPerInch / DPI
```

--- already done

Now, implement the use of opentype with an embedded font. Choose a simple open source font to embed. Show me a plan to do this and to update Game.DrawScore to use opentype font rendering.

---
