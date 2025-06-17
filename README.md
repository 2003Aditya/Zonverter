# 🔄 Zonverter

> A universal CLI file converter — convert any file to any supported format with a single command.

Zonverter is a powerful, extensible command-line tool that converts between multiple file types: JSON ↔ YAML, JPG ↔ PNG, DOCX → PDF, PPTX → DOCX, MP3 ↔ WAV, and many more.

---

## ⚙️ Features

- ✅ Supports multiple formats: text, images, audio, documents
- 🔁 One-liner conversion between any supported formats
- 🗂️ Batch conversion of entire folders
- 🧠 Smart format detection & auto-dispatcher
- 🔌 Plugin-style architecture: easy to add new converters
- 🖥️ Optional web API & GUI coming soon

---

## 📦 Installation

### Prerequisites

- [Go 1.21+](https://go.dev/)
- [`ffmpeg`](https://ffmpeg.org/download.html) (for audio conversion)
- [`LibreOffice`](https://www.libreoffice.org/download/) (for document conversion)

### Install

```bash
git clone https://github.com/yourusername/zonverter.git
cd zonverter
go build -o zonverter

