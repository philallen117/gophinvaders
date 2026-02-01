# Task Documents

This directory contains task specifications that serve as prompts for AI coding agents.

## Conventions

In the task descriptions, rectangle dimensions are give as width in pixels by height in pixels. Speeds are pixels per frame.

The file naming convention is `yyyy-mm-dd-descriptive-string.md`.

Tasks should be broken into stages that are each simple enough to for the agent to tackle as a single turn. It is OK to assume that earlier parts of the task are still in the model's context window.

When referring to file, use paths relative to the project root, without a leading slash.

When returning to a task with additional instructions, use a horizontal rule, which is three dashes on a line by itself.

Here are a snippet to help with converting Zig game config to Python game config.

> Here are some configuration constants expressed in Python. Translate them into Go and append them to cmd/gophinvaders/config.go following the same convention.
> Here are some configurations expressed in Python that all apply to THING. Translate them into Python, append them to cmd/gophinvaders/config.go following the same convention, naming them name them `thing...`.

## Purpose

These documents describe features, implementations, and changes that were requested during the development of gophinvaders. They are written in a style suitable for providing context and direction to AI assistants like GitHub Copilot.

## Status at Release

At each release, all task documents are considered completed. However, the documents themselves are not guaranteed to be accurate reflections of the final implementation, as the actual code often evolves through conversational refinement during development.

## Usage

These documents serve as:

- Historical record of feature requests and development direction.
- Reference material for understanding design decisions.
- Potential starting points for future enhancements or related features.

They should be viewed as development artifacts rather than authoritative technical documentation. For accurate information about pathman's behavior and features, refer to the README.md, user documentation, and the source code itself.
