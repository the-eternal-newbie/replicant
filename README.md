# 🧠 REPLICANT // Alpha Nexus

_"More human than human" – Tyrell Corporation_

Replicant is an offline, intelligent CLI chatbot inspired by **Blade Runner**. Powered by [Ollama](https://ollama.com) and written in **Go**, it mimics a real person’s personality and knowledge using local LLM inference (e.g., Mistral or LLaMA).

🛰 Designed to be **portable**, **trainable**, and **low-dependency**, this is the Nexus Alpha release: first in the line of human-like conversational agents.

---

## ✨ Features

- 🧬 **Personality Training**: Answer a few questions to define your replicant.
- 💬 **Conversational CLI**: Interact naturally via terminal prompts.
- 🧠 **Offline LLM Support**: Uses Ollama with models like `mistral` locally.
- 📦 **Portable & Dockerized**: Run anywhere with Docker.
- 🧩 **Modular Go Codebase**: Easy to extend and integrate.
- 🔐 **Open Source**: No API keys, no telemetry, no BS.

---

## 🛠️ Requirements

- [Docker](https://www.docker.com/) + [Docker Compose](https://docs.docker.com/compose/)
- [Ollama](https://ollama.com/) model support (e.g. `mistral`, `llama3`)

---

## 🚀 Getting Started

### 🔄 1. Clone the Project

```bash
git clone https://github.com/your-org/replicant.git
cd replicant
````

### ⚙️ 2. Configure Environment

Create a `.env` file:

```env
# .env
OLLAMA_HOST=http://ollama:11434
OLLAMA_MODEL=mistral
REPLICANT_DB=data/memory.db
```

---

### 🧪 3. Run Setup

```bash
chmod +x setup.sh
./setup.sh
```

This will:

* Stop any old containers
* Build and start the Replicant + Ollama services
* Auto-download the `mistral` model
* Launch your chatbot

---

## 🧬 Training Your Replicant

When you first launch, you'll be prompted to enter:

* Name
* Skills
* Tone (e.g. sarcastic, professional)
* Interests
* Fun facts
* Known Q\&A (add your own knowledge manually)

These values are saved into a local BoltDB file (`memory.db`) and used to generate a **system prompt** that instructs the LLM how to behave.

---

## 💬 Talking to Your Replicant

Once trained, choose “chat” mode and start interacting:

```bash
docker attach replicant
```

Or run manually:

```bash
docker exec -it replicant /usr/local/bin/replicant
```

Type questions in English and receive context-aware, personalized answers.

---

## 🧰 Development Notes

* All code is in Go (`main.go` + modular files in `/replicant`)
* LLM communication via `http://ollama:11434/api/generate`
* BoltDB used for persistent knowledge and personality
* `.env` file supports host/model/configurable DB path

---

## 🧠 Roadmap

* [ ] Web UI with terminal-style interface
* [ ] Contextual memory (short-term conversation)
* [ ] Knowledge embedding / vector DB support
* [ ] Multiple profiles / Replicant switching
* [ ] Nexus Beta release

---

## 📂 Project Structure

```
replicant/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── .env
├── setup.sh
├── data/
│   └── memory.db
├── replicant/
│   ├── engine.go
│   ├── ollama.go
│   ├── personality.go
│   └── memory.go
```

---

## 🤖 Version

**Nexus Alpha** – *First release candidate of the Replicant line*
This version is CLI-only, deterministic, and locally intelligent.

---

## ☂️ License

This project is released under the [MIT License](LICENSE).
Built by engineers. Not by Tyrell Corp.

---

> “The light that burns twice as bright burns half as long. And you have burned so very, very brightly.” – Eldon Tyrell

