# Project: North Star

## Description
A CLI tool to keep track of your goals and activities or checkpoints related to them.

## MVP
Note-taking in the CLI

## Working Principle
```
Model (state) → re-render → View (UI) → key pressed (message) → Update (logic) → update state → Model (state)
```

**Architecture Flow:**
- **Model** maintains application state
- **View** renders the current state to UI
- **Update** processes user input messages and updates state
- Loop continues with re-rendering