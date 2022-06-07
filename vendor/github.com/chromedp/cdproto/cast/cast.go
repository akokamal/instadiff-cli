// Package cast provides the Chrome DevTools Protocol
// commands, types, and events for the Cast domain.
//
// A domain for interacting with Cast, Presentation API, and Remote Playback
// API functionalities.
//
// Generated by the cdproto-gen command.
package cast

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// EnableParams starts observing for sinks that can be used for tab
// mirroring, and if set, sinks compatible with |presentationUrl| as well. When
// sinks are found, a |sinksUpdated| event is fired. Also starts observing for
// issue messages. When an issue is added or removed, an |issueUpdated| event is
// fired.
type EnableParams struct {
	PresentationURL string `json:"presentationUrl,omitempty"`
}

// Enable starts observing for sinks that can be used for tab mirroring, and
// if set, sinks compatible with |presentationUrl| as well. When sinks are
// found, a |sinksUpdated| event is fired. Also starts observing for issue
// messages. When an issue is added or removed, an |issueUpdated| event is
// fired.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-enable
//
// parameters:
func Enable() *EnableParams {
	return &EnableParams{}
}

// WithPresentationURL [no description].
func (p EnableParams) WithPresentationURL(presentationURL string) *EnableParams {
	p.PresentationURL = presentationURL
	return &p
}

// Do executes Cast.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, p, nil)
}

// DisableParams stops observing for sinks and issues.
type DisableParams struct{}

// Disable stops observing for sinks and issues.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Cast.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// SetSinkToUseParams sets a sink to be used when the web page requests the
// browser to choose a sink via Presentation API, Remote Playback API, or Cast
// SDK.
type SetSinkToUseParams struct {
	SinkName string `json:"sinkName"`
}

// SetSinkToUse sets a sink to be used when the web page requests the browser
// to choose a sink via Presentation API, Remote Playback API, or Cast SDK.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-setSinkToUse
//
// parameters:
//   sinkName
func SetSinkToUse(sinkName string) *SetSinkToUseParams {
	return &SetSinkToUseParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.setSinkToUse against the provided context.
func (p *SetSinkToUseParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetSinkToUse, p, nil)
}

// StartDesktopMirroringParams starts mirroring the desktop to the sink.
type StartDesktopMirroringParams struct {
	SinkName string `json:"sinkName"`
}

// StartDesktopMirroring starts mirroring the desktop to the sink.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-startDesktopMirroring
//
// parameters:
//   sinkName
func StartDesktopMirroring(sinkName string) *StartDesktopMirroringParams {
	return &StartDesktopMirroringParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.startDesktopMirroring against the provided context.
func (p *StartDesktopMirroringParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartDesktopMirroring, p, nil)
}

// StartTabMirroringParams starts mirroring the tab to the sink.
type StartTabMirroringParams struct {
	SinkName string `json:"sinkName"`
}

// StartTabMirroring starts mirroring the tab to the sink.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-startTabMirroring
//
// parameters:
//   sinkName
func StartTabMirroring(sinkName string) *StartTabMirroringParams {
	return &StartTabMirroringParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.startTabMirroring against the provided context.
func (p *StartTabMirroringParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartTabMirroring, p, nil)
}

// StopCastingParams stops the active Cast session on the sink.
type StopCastingParams struct {
	SinkName string `json:"sinkName"`
}

// StopCasting stops the active Cast session on the sink.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#method-stopCasting
//
// parameters:
//   sinkName
func StopCasting(sinkName string) *StopCastingParams {
	return &StopCastingParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.stopCasting against the provided context.
func (p *StopCastingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStopCasting, p, nil)
}

// Command names.
const (
	CommandEnable                = "Cast.enable"
	CommandDisable               = "Cast.disable"
	CommandSetSinkToUse          = "Cast.setSinkToUse"
	CommandStartDesktopMirroring = "Cast.startDesktopMirroring"
	CommandStartTabMirroring     = "Cast.startTabMirroring"
	CommandStopCasting           = "Cast.stopCasting"
)
