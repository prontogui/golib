# ProntoGUI

**The FAST way to build a GUI — in Go.**

Build a desktop GUI for your Go service or program without learning React, Flutter, Qt, or any other frontend stack. Write your UI directly in Go and let the ProntoGUI App render it on macOS or Windows.

Website: [www.prontogui.com](https://www.prontogui.com)

<!-- TODO: replace with a real demo GIF (e.g. docs/demo.gif). Keep it short (~5s) and showcase a live UI updating in response to a button press. -->
<p align="center">
  <img src="https://prontogui.com/wp-content/uploads/2026/05/pg_demo_clip.gif" alt="ProntoGUI demo" width="640">
</p>

[![Go Reference](https://pkg.go.dev/badge/github.com/prontogui/golib.svg)](https://pkg.go.dev/github.com/prontogui/golib)
[![License: BSD-3-Clause](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](LICENSE)

---

## Why ProntoGUI?

Modern GUI development is fragmented across React, React Native, .NET, Qt, Flutter, and more. If you write Go, shipping a GUI usually means either learning a whole frontend stack or waiting on a frontend developer.

ProntoGUI removes that barrier:

- **Stay in Go.** Define your GUI with plain Go types — no JavaScript, no XAML, no QML.
- **Skip the frontend.** A pre-built native App renders your GUI. You ship a single Go binary; users install the App once.
- **Streaming, not request/response.** The GUI is live — set primitives, wait for events, update state, repeat.

Ideal for internal tools, manufacturing/lab apps, ERP front-ends, simulators, and any Go service that needs a visual front-end without the frontend team.

## Install

```sh
go get github.com/prontogui/golib
```

Requires Go 1.23+.

## Quickstart

**1. Install the ProntoGUI App.** Your Go program is the server; the App is the renderer.
App is available for macOS and Windows through the early access program with an official launch soon **[ProntoGUI Early Access Request](https://www.prontogui.com/early-access)**.

**2. Write your program:**

```go
package main

import (
    "log"

    "github.com/prontogui/golib"
)

func main() {
    pgui := golib.NewProntoGUI()
    if err := pgui.StartServingSingle("127.0.0.1", 50053); err != nil {
        log.Fatal(err)
    }

    greeting := golib.NewText("Hello from Go!")
    button := golib.CommandWith{Label: "Click me"}.Make()

    pgui.SetGUI(greeting, button)

    for {
        updated, err := pgui.Wait()
        if err != nil {
            log.Fatal(err)
        }
        if updated == button {
            greeting.SetContent("Button clicked!")
        }
    }
}
```

**3. Run it, then open the App** and point it at `127.0.0.1:50053`.

## The App

ProntoGUI ships with a native rendering App built on Flutter. Think of it as a browser for ProntoGUI servers — your Go program emits primitives over gRPC, the App draws them.

- Installers for **macOS, Windows** (Intel, Apple Silicon, ARM)
- Fast launch, compact binary
- Regularly updated — security and bug fixes handled by us

**Get the App at **[ProntoGUI Early Access Request](https://www.prontogui.com/early-access)****

## Features

- 20+ primitives: `Text`, `Command`, `Check`, `Choice`, `TextField`, `NumericField`, `Table`, `List`, `Frame`, `Group`, `Card`, `Image`, `Icon`, `Timer`, `ImportFile`, `ExportFile`, and more
- 8,000+ built-in icons
- Flow, pixel-positioning, and box-model layouts
- International (Unicode) text support
- Single-client and multi-client server modes
- gRPC over HTTP/2 wire protocol — secure, efficient, language-agnostic

## Documentation

- **API reference:** [pkg.go.dev/github.com/prontogui/golib](https://pkg.go.dev/github.com/prontogui/golib)
- **Introduction & tutorial:** [prontogui.com/support/intro](https://prontogui.com/support/intro)
- **Primitives reference:** [prontogui.com/support/ref-guide](https://prontogui.com/support/ref-guide/)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Issues and pull requests welcome.

## License

BSD 3-Clause — see [LICENSE](LICENSE).

---

###### Copyright 2024–2026 ProntoGUI, LLC
###### ProntoGUI™ is a trademark of ProntoGUI, LLC
