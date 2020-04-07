package astilectron

import (
	"context"
	"fmt"
	stdUrl "net/url"
	"path/filepath"
	"sync"

	astikit "github.com/asticode/go-astikit"
)

// Window event names
const (
	EventNameWebContentsEventLogin             = "web.contents.event.login"
	EventNameWebContentsEventLoginCallback     = "web.contents.event.login.callback"
	EventNameWindowCmdBlur                     = "window.cmd.blur"
	EventNameWindowCmdCenter                   = "window.cmd.center"
	EventNameWindowCmdClose                    = "window.cmd.close"
	EventNameWindowCmdCreate                   = "window.cmd.create"
	EventNameWindowCmdCreateBrowserView        = "window.cmd.createbrowserview"
	EventNameWindowCmdDestroy                  = "window.cmd.destroy"
	EventNameWindowCmdFocus                    = "window.cmd.focus"
	EventNameWindowCmdHide                     = "window.cmd.hide"
	EventNameWindowCmdLog                      = "window.cmd.log"
	EventNameWindowCmdMaximize                 = "window.cmd.maximize"
	EventNameWindowCmdIsMaximized              = "window.cmd.ismaximized"
	EventNameWindowCmdSetFullscreen            = "window.cmd.setfullscreen"
	EventNameWindowCmdSetFrameShow             = "window.cmd.setframeshow"
	eventNameWindowCmdMessage                  = "window.cmd.message"
	eventNameWindowCmdMessageCallback          = "window.cmd.message.callback"
	EventNameWindowCmdMinimize                 = "window.cmd.minimize"
	EventNameWindowCmdMove                     = "window.cmd.move"
	EventNameWindowCmdResize                   = "window.cmd.resize"
	EventNameWindowCmdSetAutoResize            = "window.cmd.setautoresize"
	EventNameWindowCmdGetBounds                = "window.cmd.getbounds"
	EventNameWindowCmdSetBounds                = "window.cmd.set.bounds"
	EventNameWindowCmdRestore                  = "window.cmd.restore"
	EventNameWindowCmdShow                     = "window.cmd.show"
	EventNameWindowCmdUnmaximize               = "window.cmd.unmaximize"
	EventNameWindowCmdWebContentsCloseDevTools = "window.cmd.web.contents.close.dev.tools"
	EventNameWindowCmdWebContentsOpenDevTools  = "window.cmd.web.contents.open.dev.tools"
	EventNameWindowEventBlur                   = "window.event.blur"
	EventNameWindowEventClosed                 = "window.event.closed"
	EventNameWindowEventDidFinishLoad          = "window.event.did.finish.load"
	EventNameWindowEventFocus                  = "window.event.focus"
	EventNameWindowEventHide                   = "window.event.hide"
	EventNameWindowEventMaximize               = "window.event.maximize"
	EventNameWindowEventIsMaximized            = "window.event.ismaximized"
	EventNameWindowEventSetFullscreen          = "window.event.setfullscreen"
	EventNameWindowEventSetFrameShow           = "window.event.setframeshow"
	eventNameWindowEventMessage                = "window.event.message"
	eventNameWindowEventMessageCallback        = "window.event.message.callback"
	EventNameWindowEventMinimize               = "window.event.minimize"
	EventNameWindowEventMove                   = "window.event.move"
	EventNameWindowEventReadyToShow            = "window.event.ready.to.show"
	EventNameWindowEventResize                 = "window.event.resize"
	EventNameWindowEventSetBounds              = "window.event.setbounds"
	EventNameWindowEventSetAutoResize          = "window.event.setautoresize"
	EventNameWindowEventGetBounds              = "window.event.getbounds"
	EventNameWindowEventRestore                = "window.event.restore"
	EventNameWindowEventShow                   = "window.event.show"
	EventNameWindowEventUnmaximize             = "window.event.unmaximize"
	EventNameWindowEventUnresponsive           = "window.event.unresponsive"
	EventNameWindowEventDidGetRedirectRequest  = "window.event.did.get.redirect.request"
	EventNameWindowEventWillNavigate           = "window.event.will.navigate"
	EventNameWindowEventNewWindow              = "window.event.newwindow"
)

// Title bar styles
var (
	TitleBarStyleDefault     = astikit.StrPtr("default")
	TitleBarStyleHidden      = astikit.StrPtr("hidden")
	TitleBarStyleHiddenInset = astikit.StrPtr("hidden-inset")
)

// Window represents a window
// TODO Add missing window options
// TODO Add missing window methods
// TODO Add missing window events
type Window struct {
	*object
	callbackIdentifier *identifier
	l                  astikit.SeverityLogger
	m                  sync.Mutex // Locks o
	o                  *WindowOptions
	onMessageOnce      sync.Once
	Session            *Session
	url                *stdUrl.URL
}

// WindowOptions represents window options
// We must use pointers since GO doesn't handle optional fields whereas NodeJS does. Use astikit.BoolPtr, astikit.IntPtr or astikit.StrPtr
// to fill the struct
// https://github.com/electron/electron/blob/v1.8.1/docs/api/browser-window.md
type WindowOptions struct {
	AcceptFirstMouse       *bool           `json:"acceptFirstMouse,omitempty"`
	AlwaysOnTop            *bool           `json:"alwaysOnTop,omitempty"`
	AutoHideMenuBar        *bool           `json:"autoHideMenuBar,omitempty"`
	BackgroundColor        *string         `json:"backgroundColor,omitempty"`
	Center                 *bool           `json:"center,omitempty"`
	Closable               *bool           `json:"closable,omitempty"`
	DisableAutoHideCursor  *bool           `json:"disableAutoHideCursor,omitempty"`
	EnableLargerThanScreen *bool           `json:"enableLargerThanScreen,omitempty"`
	Focusable              *bool           `json:"focusable,omitempty"`
	Frame                  *bool           `json:"frame,omitempty"`
	Fullscreen             *bool           `json:"fullscreen,omitempty"`
	Fullscreenable         *bool           `json:"fullscreenable,omitempty"`
	HasShadow              *bool           `json:"hasShadow,omitempty"`
	Height                 *int            `json:"height,omitempty"`
	Icon                   *string         `json:"icon,omitempty"`
	Kiosk                  *bool           `json:"kiosk,omitempty"`
	MaxHeight              *int            `json:"maxHeight,omitempty"`
	Maximizable            *bool           `json:"maximizable,omitempty"`
	Maximized              *bool           `json:"maximized,omitempty"`
	MaxWidth               *int            `json:"maxWidth,omitempty"`
	MinHeight              *int            `json:"minHeight,omitempty"`
	Minimizable            *bool           `json:"minimizable,omitempty"`
	MinWidth               *int            `json:"minWidth,omitempty"`
	Modal                  *bool           `json:"modal,omitempty"`
	Movable                *bool           `json:"movable,omitempty"`
	Resizable              *bool           `json:"resizable,omitempty"`
	Show                   *bool           `json:"show,omitempty"`
	SkipTaskbar            *bool           `json:"skipTaskbar,omitempty"`
	Title                  *string         `json:"title,omitempty"`
	TitleBarStyle          *string         `json:"titleBarStyle,omitempty"`
	Transparent            *bool           `json:"transparent,omitempty"`
	UseContentSize         *bool           `json:"useContentSize,omitempty"`
	WebPreferences         *WebPreferences `json:"webPreferences,omitempty"`
	Width                  *int            `json:"width,omitempty"`
	X                      *int            `json:"x,omitempty"`
	Y                      *int            `json:"y,omitempty"`
	BrowserView            *bool           `json:"browserView,omitempty"`
	ParentID               *string         `json:"parentID,omitempty"`

	// Additional options
	AppDetails *WindowAppDetails    `json:"appDetails,omitempty"`
	Custom     *WindowCustomOptions `json:"custom,omitempty"`
	Load       *WindowLoadOptions   `json:"load,omitempty"`
	Proxy      *WindowProxyOptions  `json:"proxy,omitempty"`
}

// WindowAppDetails represents window app details
// https://github.com/electron/electron/blob/v4.0.1/docs/api/browser-window.md#winsetappdetailsoptions-windows
type WindowAppDetails struct {
	AppID               *string `json:"appId,omitempty"`
	AppIconPath         *string `json:"appIconPath,omitempty"`
	RelaunchCommand     *string `json:"relaunchCommand,omitempty"`
	AppIconIndex        *int    `json:"appIconIndex,omitempty"`
	RelaunchDisplayName *string `json:"relaunchDisplayName,omitempty"`
}

// WindowCustomOptions represents window custom options
type WindowCustomOptions struct {
	HideOnClose       *bool              `json:"hideOnClose,omitempty"`
	MessageBoxOnClose *MessageBoxOptions `json:"messageBoxOnClose,omitempty"`
	MinimizeOnClose   *bool              `json:"minimizeOnClose,omitempty"`
	Script            string             `json:"script,omitempty"`
}

// WindowLoadOptions represents window load options
// https://github.com/electron/electron/blob/v1.8.1/docs/api/browser-window.md#winloadurlurl-options
type WindowLoadOptions struct {
	ExtraHeaders string `json:"extraHeaders,omitempty"`
	HTTPReferer  string `json:"httpReferrer,omitempty"`
	UserAgent    string `json:"userAgent,omitempty"`
}

// WindowProxyOptions represents window proxy options
// https://github.com/electron/electron/blob/v1.8.1/docs/api/session.md#sessetproxyconfig-callback
type WindowProxyOptions struct {
	BypassRules string `json:"proxyBypassRules,omitempty"`
	PACScript   string `json:"pacScript,omitempty"`
	Rules       string `json:"proxyRules,omitempty"`
}

// WebPreferences represents web preferences in window options.
// We must use pointers since GO doesn't handle optional fields whereas NodeJS does.
// Use astikit.BoolPtr, astikit.IntPtr or astikit.StrPtr to fill the struct
type WebPreferences struct {
	AllowRunningInsecureContent *bool                  `json:"allowRunningInsecureContent,omitempty"`
	BackgroundThrottling        *bool                  `json:"backgroundThrottling,omitempty"`
	BlinkFeatures               *string                `json:"blinkFeatures,omitempty"`
	ContextIsolation            *bool                  `json:"contextIsolation,omitempty"`
	DefaultEncoding             *string                `json:"defaultEncoding,omitempty"`
	DefaultFontFamily           map[string]interface{} `json:"defaultFontFamily,omitempty"`
	DefaultFontSize             *int                   `json:"defaultFontSize,omitempty"`
	DefaultMonospaceFontSize    *int                   `json:"defaultMonospaceFontSize,omitempty"`
	DevTools                    *bool                  `json:"devTools,omitempty"`
	DisableBlinkFeatures        *string                `json:"disableBlinkFeatures,omitempty"`
	ExperimentalCanvasFeatures  *bool                  `json:"experimentalCanvasFeatures,omitempty"`
	ExperimentalFeatures        *bool                  `json:"experimentalFeatures,omitempty"`
	Images                      *bool                  `json:"images,omitempty"`
	Javascript                  *bool                  `json:"javascript,omitempty"`
	MinimumFontSize             *int                   `json:"minimumFontSize,omitempty"`
	// This attribute needs to be true at all time
	// NodeIntegration             *bool                  `json:"nodeIntegration,omitempty"`
	NodeIntegrationInWorker *bool                  `json:"nodeIntegrationInWorker,omitempty"`
	Offscreen               *bool                  `json:"offscreen,omitempty"`
	Partition               *string                `json:"partition,omitempty"`
	Plugins                 *bool                  `json:"plugins,omitempty"`
	Preload                 *string                `json:"preload,omitempty"`
	Sandbox                 *bool                  `json:"sandbox,omitempty"`
	ScrollBounce            *bool                  `json:"scrollBounce,omitempty"`
	Session                 map[string]interface{} `json:"session,omitempty"`
	TextAreasAreResizable   *bool                  `json:"textAreasAreResizable,omitempty"`
	Webaudio                *bool                  `json:"webaudio,omitempty"`
	Webgl                   *bool                  `json:"webgl,omitempty"`
	WebSecurity             *bool                  `json:"webSecurity,omitempty"`
	WebviewTag              *bool                  `json:"webviewTag,omitempty"`
	ZoomFactor              *float64               `json:"zoomFactor,omitempty"`
}

// newWindow creates a new window
func newWindow(ctx context.Context, l astikit.SeverityLogger, o Options, p Paths, url string, wo *WindowOptions, d *dispatcher, i *identifier, wrt *writer) (w *Window, err error) {
	// Init
	w = &Window{
		callbackIdentifier: newIdentifier(),
		l:                  l,
		o:                  wo,
		object:             newObject(ctx, d, i, wrt, i.new()),
	}
	w.Session = newSession(w.ctx, d, i, wrt)

	// Check app details
	if wo.Icon == nil && p.AppIconDefaultSrc() != "" {
		wo.Icon = astikit.StrPtr(p.AppIconDefaultSrc())
	}
	if wo.Title == nil && o.AppName != "" {
		wo.Title = astikit.StrPtr(o.AppName)
	}

	// Make sure the window's context is cancelled once the closed event is received
	w.On(EventNameWindowEventClosed, func(e Event) (deleteListener bool) {
		w.cancel()
		return true
	})

	// Show
	w.On(EventNameWindowEventHide, func(e Event) (deleteListener bool) {
		w.m.Lock()
		defer w.m.Unlock()
		w.o.Show = astikit.BoolPtr(false)
		return
	})
	w.On(EventNameWindowEventShow, func(e Event) (deleteListener bool) {
		w.m.Lock()
		defer w.m.Unlock()
		w.o.Show = astikit.BoolPtr(true)
		return
	})
	// Basic parse
	if w.url, err = stdUrl.Parse(url); err != nil {
		err = fmt.Errorf("std parsing of url %s failed: %w", url, err)
		return
	}

	// File
	if w.url.Scheme == "" {
		// Get absolute path
		if url, err = filepath.Abs(url); err != nil {
			err = fmt.Errorf("getting absolute path of %s failed: %w", url, err)
			return
		}

		// Set url
		w.url = &stdUrl.URL{Path: filepath.ToSlash(url), Scheme: "file"}
	}
	return
}

// NewMenu creates a new window menu
func (w *Window) NewMenu(i []*MenuItemOptions) *Menu {
	return newMenu(w.ctx, w.id, i, w.d, w.i, w.w)
}

// Blur blurs the window
func (w *Window) Blur() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdBlur, TargetID: w.id}, EventNameWindowEventBlur)
	return
}

// Center centers the window
func (w *Window) Center() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdCenter, TargetID: w.id}, EventNameWindowEventMove)
	return
}

// Close closes the window
func (w *Window) Close() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdClose, TargetID: w.id}, EventNameWindowEventClosed)
	return
}

// CloseDevTools closes the dev tools
func (w *Window) CloseDevTools() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	return w.w.write(Event{Name: EventNameWindowCmdWebContentsCloseDevTools, TargetID: w.id})
}

// Create creates the window
// We wait for EventNameWindowEventDidFinishLoad since we need the web content to be fully loaded before being able to
// send messages to it
func (w *Window) Create() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdCreate, SessionID: w.Session.id, TargetID: w.id, URL: w.url.String(), WindowOptions: w.o}, EventNameWindowEventDidFinishLoad)
	return
}

// Create creates the browser window
// We wait for EventNameWindowEventDidFinishLoad since we need the web content to be fully loaded before being able to
// send messages to it
func (w *Window) CreateBrowserView(pw *Window) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	if pw == nil {
		err = fmt.Errorf("parent window has not to be nil")
		return
	}
	// set parent id of this browser view
	w.o.ParentID = astikit.StrPtr(pw.id)
	w.o.BrowserView = astikit.BoolPtr(true)

	_, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:          EventNameWindowCmdCreateBrowserView,
			SessionID:     w.Session.id,
			TargetID:      w.id,
			URL:           w.url.String(),
			WindowOptions: w.o,
		}, EventNameWindowEventDidFinishLoad)
	// set bounds
	if err == nil {
		w.SetBounds(RectangleOptions{
			PositionOptions: PositionOptions{
				X: w.o.X,
				Y: w.o.Y,
			},
			SizeOptions: SizeOptions{
				Height: w.o.Height,
				Width:  w.o.Width,
			},
		})
	}
	return
}

// Destroy destroys the window
func (w *Window) Destroy() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdDestroy, TargetID: w.id}, EventNameWindowEventClosed)
	return
}

// Focus focuses on the window
func (w *Window) Focus() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdFocus, TargetID: w.id}, EventNameWindowEventFocus)
	return
}

// Hide hides the window
func (w *Window) Hide() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdHide, TargetID: w.id}, EventNameWindowEventHide)
	return
}

// IsShown returns whether the window is shown
func (w *Window) IsShown() bool {
	if w.ctx.Err() != nil {
		return false
	}
	w.m.Lock()
	defer w.m.Unlock()
	return w.o.Show != nil && *w.o.Show
}

// Log logs a message in the JS console of the window
func (w *Window) Log(message string) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	return w.w.write(Event{Message: newEventMessage(message), Name: EventNameWindowCmdLog, TargetID: w.id})
}

// Maximize maximizes the window
func (w *Window) Maximize() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdMaximize, TargetID: w.id}, EventNameWindowEventMaximize)
	return
}

// check the window is maximized
func (w *Window) IsMaximized() (is bool, err error) {
	if err = w.ctx.Err(); err != nil {
		is = false
		return
	}
	var be Event
	be, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:     EventNameWindowCmdIsMaximized,
			TargetID: w.id,
		},
		EventNameWindowEventIsMaximized)
	if err == nil && be.WindowOptions != nil {
		is = *be.WindowOptions.Maximized
	}
	return
}

// fullscreen the window
func (w *Window) SetFullScreen(fullscreen bool) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	w.m.Lock()
	w.o.Fullscreen = astikit.BoolPtr(fullscreen)
	_, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:     EventNameWindowCmdSetFullscreen,
			TargetID: w.id,
			WindowOptions: &WindowOptions{
				Fullscreen: w.o.Fullscreen,
			},
		}, EventNameWindowEventSetFullscreen)
	return
}

// set frame show or not
func (w *Window) SetFrameShow(showframe bool) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	w.m.Lock()
	w.o.Frame = astikit.BoolPtr(showframe)
	_, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:     EventNameWindowCmdSetFrameShow,
			TargetID: w.id,
			WindowOptions: &WindowOptions{
				Frame: w.o.Frame,
			},
		}, EventNameWindowEventSetFrameShow)
	return
}

// Minimize minimizes the window
func (w *Window) Minimize() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdMinimize, TargetID: w.id}, EventNameWindowEventMinimize)
	return
}

// Move moves the window
func (w *Window) Move(x, y int) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	w.m.Lock()
	w.o.X = astikit.IntPtr(x)
	w.o.Y = astikit.IntPtr(y)
	w.m.Unlock()
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdMove, TargetID: w.id, WindowOptions: &WindowOptions{X: astikit.IntPtr(x), Y: astikit.IntPtr(y)}}, EventNameWindowEventMove)
	return
}

// MoveInDisplay moves the window in the proper display
func (w *Window) MoveInDisplay(d *Display, x, y int) error {
	return w.Move(d.Bounds().X+x, d.Bounds().Y+y)
}

func (w *Window) OnLogin(fn func(i Event) (username, password string, err error)) {
	w.On(EventNameWebContentsEventLogin, func(i Event) (deleteListener bool) {
		// Get username and password
		username, password, err := fn(i)
		if err != nil {
			w.l.Error(fmt.Errorf("getting username and password failed: %w", err))
			return
		}

		// No auth
		if len(username) == 0 && len(password) == 0 {
			return
		}

		// Send message back
		if err = w.w.write(Event{CallbackID: i.CallbackID, Name: EventNameWebContentsEventLoginCallback, Password: password, TargetID: w.id, Username: username}); err != nil {
			w.l.Error(fmt.Errorf("writing login callback message failed: %w", err))
			return
		}
		return
	})
}

// ListenerMessage represents a message listener executed when receiving a message from the JS
type ListenerMessage func(m *EventMessage) (v interface{})

// OnMessage adds a specific listener executed when receiving a message from the JS
// This method can be called only once
func (w *Window) OnMessage(l ListenerMessage) {
	w.onMessageOnce.Do(func() {
		w.On(eventNameWindowEventMessage, func(i Event) (deleteListener bool) {
			v := l(i.Message)
			if len(i.CallbackID) > 0 {
				o := Event{CallbackID: i.CallbackID, Name: eventNameWindowCmdMessageCallback, TargetID: w.id}
				if v != nil {
					o.Message = newEventMessage(v)
				}
				if err := w.w.write(o); err != nil {
					w.l.Error(fmt.Errorf("writing callback message failed: %w", err))
				}
			}
			return
		})
	})
}

// ListenerMessage represents a message listener executed when receiving a message from the JS
type ListenerBounds func(m *RectangleOptions) (v interface{})

// OnMove event
func (w *Window) OnMove(l ListenerBounds) {
	w.On(EventNameWindowEventMove, func(i Event) (deleteListener bool) {
		l(i.Bounds)
		return
	})
}

// OpenDevTools opens the dev tools
func (w *Window) OpenDevTools() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	return w.w.write(Event{Name: EventNameWindowCmdWebContentsOpenDevTools, TargetID: w.id})
}

// Resize resizes the window
func (w *Window) Resize(width, height int) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	w.m.Lock()
	w.o.Height = astikit.IntPtr(height)
	w.o.Width = astikit.IntPtr(width)
	w.m.Unlock()
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdResize, TargetID: w.id, WindowOptions: &WindowOptions{Height: astikit.IntPtr(height), Width: astikit.IntPtr(width)}}, EventNameWindowEventResize)
	return
}

// SetBounds set bounds of the window
func (w *Window) SetBounds(r RectangleOptions) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	w.m.Lock()
	w.o.Height = r.Height
	w.o.Width = r.Width
	w.o.X = r.X
	w.o.Y = r.Y
	w.m.Unlock()
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdSetBounds, TargetID: w.id, Bounds: &r}, EventNameWindowEventResize)
	return
}

// Restore restores the window
func (w *Window) Restore() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdRestore, TargetID: w.id}, EventNameWindowEventRestore)
	return
}

// CallbackMessage represents a message callback
type CallbackMessage func(m *EventMessage)

// SendMessage sends a message to the JS window and execute optional callbacks upon receiving a response from the JS
// Use astilectron.onMessage method to capture those messages in JS
func (w *Window) SendMessage(message interface{}, callbacks ...CallbackMessage) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	var e = Event{Message: newEventMessage(message), Name: eventNameWindowCmdMessage, TargetID: w.id}
	if len(callbacks) > 0 {
		e.CallbackID = w.callbackIdentifier.new()
		w.On(eventNameWindowEventMessageCallback, func(i Event) (deleteListener bool) {
			if i.CallbackID == e.CallbackID {
				for _, c := range callbacks {
					c(i.Message)
				}
				deleteListener = true
			}
			return
		})
	}
	return w.w.write(e)
}

// Show shows the window
func (w *Window) Show() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdShow, TargetID: w.id}, EventNameWindowEventShow)
	return
}

// Unmaximize unmaximize the window
func (w *Window) Unmaximize() (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w, Event{Name: EventNameWindowCmdUnmaximize, TargetID: w.id}, EventNameWindowEventUnmaximize)
	return
}

/*
// setBounds the window or browserview
func (w *Window) SetBounds(bounds *RectangleOptions) (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	_, err = synchronousEvent(w.c, w, w.w,
		Event{
			Name:     EventNameWindowCmdSetBounds,
			TargetID: w.id,
			Bounds:   bounds,
		}, EventNameWindowEventSetBounds)
	return
}
*/

// getBounds of the window or browserview
func (w *Window) GetBounds() (bounds *RectangleOptions, err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	var be Event
	be, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:     EventNameWindowCmdGetBounds,
			TargetID: w.id,
		}, EventNameWindowEventGetBounds)
	if err == nil && be.Bounds != nil {
		bounds = be.Bounds
	}
	return
}

// browser view setAutoResize
func (w *Window) SetAutoResize(aResize *AutoResizeOptions) (err error) {
	if err = w.ctx.Err(); err != nil {
		return
	}
	_, err = synchronousEvent(w.ctx, w, w.w,
		Event{
			Name:              EventNameWindowCmdSetAutoResize,
			TargetID:          w.id,
			AutoResizeOptions: aResize,
		}, EventNameWindowEventSetAutoResize)
	return
}

//
type ListenerEvent func(data interface{}) (v interface{})

// On event
func (w *Window) OnEvent(eventName string, l ListenerEvent, once bool) {
	w.On(eventName, func(i Event) (deleteListener bool) {
		l(i)
		deleteListener = once
		return
	})
}

// OnNewWindow event
func (w *Window) OnNewWindow(l ListenerEvent) {
	w.OnEvent(EventNameWindowEventNewWindow, l, false)
}

// OnReadyToShow event
func (w *Window) OnReadyToShow(l ListenerEvent) {
	w.OnEvent(EventNameWindowEventReadyToShow, l, false)
}
