```golang
type DispatchMouseEventParams struct {
	Type               MouseType                     `json:"type"`                         // Type of the mouse event.
	X                  float64                       `json:"x"`                            // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y                  float64                       `json:"y"`                            // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Modifiers          Modifier                      `json:"modifiers"`                    // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp          *TimeSinceEpoch               `json:"timestamp,omitempty"`          // Time at which the event occurred.
	Button             MouseButton                   `json:"button,omitempty"`             // Mouse button (default: "none").
	Buttons            int64                         `json:"buttons,omitempty"`            // A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
	ClickCount         int64                         `json:"clickCount,omitempty"`         // Number of times the mouse button was clicked (default: 0).
	Force              float64                       `json:"force,omitempty"`              // The normalized pressure, which has a range of [0,1] (default: 0).
	TangentialPressure float64                       `json:"tangentialPressure,omitempty"` // The normalized tangential pressure, which has a range of [-1,1] (default: 0).
	TiltX              int64                         `json:"tiltX,omitempty"`              // The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0).
	TiltY              int64                         `json:"tiltY,omitempty"`              // The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	Twist              int64                         `json:"twist,omitempty"`              // The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
	DeltaX             float64                       `json:"deltaX"`                       // X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY             float64                       `json:"deltaY"`                       // Y delta in CSS pixels for mouse wheel event (default: 0).
	PointerType        DispatchMouseEventPointerType `json:"pointerType,omitempty"`        // Pointer type (default: "mouse").
}
```

MouseType: 按下，释放，移动，滚轮
```golang
const (
	MousePressed  MouseType = "mousePressed"
	MouseReleased MouseType = "mouseReleased"
	MouseMoved    MouseType = "mouseMoved"
	MouseWheel    MouseType = "mouseWheel"
)
```

MouseButton: none，左键，中间键，右键，后退键，前进键。多数只有三个吧（左键，中间键(可以滚，可以按)，右键）
```golang
const (
	None    MouseButton = "none"
	Left    MouseButton = "left"
	Middle  MouseButton = "middle"
	Right   MouseButton = "right"
	Back    MouseButton = "back"
	Forward MouseButton = "forward"
)
```

X, Y: 标明鼠标的绝对位置，点击事件的时候需要用到
DeltaX, DeltaY: 滚轮的时候主要用到，用来标明一次移动多少Delta
TiltX, TiltY: ???


如何拖动鼠标：
1. 设置 Type = MousePressed 然后do
2. 设置 Type = MouseMoved 然后不同更换 X,Y 然后DO
3. 设置 Type = MouseReleased 然后Do。完成了拖动

点击鼠标就简单了

鼠标滑轮滚动：
1. 设置 Type = MousePressed 
2. 设置 DeltaX, DeltaY, 然后不停的 Do 就可以了！

参数中有 X，Y值需要设置，那么如果想要点击某个 元素时，如何获得元素的位置呢？ `dom.GetContentQuads().WithNodeID(id).Do(ctx)`
```golang
// 找到 sel 节点后，会之后 func 里面指定的操作。
chromedp.QueryAfter(sel, func(ctx context.Context, id runtime.ExecutionContextID, node ...*cdp.Node) error {
		if len(node) == 0 {
			fmt.Println("找不到相应node")
			return fmt.Errorf("找不到相关 Node")
		}
        // 获取节点的 坐标 quads里面四个值，分别是 左上X，左上Y，右下X，右下Y
        quads, err := dom.GetContentQuads().WithNodeID(node[0].NodeID).Do(ctx)
        return err
	})
```

## 其它
获取selector所选定的所有节点的 Text

```golang
texts := make([]string, 0)
chromedp.Evaluate(`[...document.querySelectorAll('#tags li')].map((e) => e.innerText)`, &texts)
// https://github.com/chromedp/chromedp/issues/87
```
