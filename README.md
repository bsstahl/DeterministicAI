# Deterministic AI Samples

This repository is intended to become a collection of small, practical samples showing how to build deterministic AI models in different programming languages.

## What is Deterministic AI?

Here, **Deterministic AI** means using an explicit model of a problem and a
solver to choose an answer that follows defined rules and constraints. The
model describes the choices that can be made, the goal to optimize (if any),
and the limits that a valid answer must obey.

Given the same model and inputs, these methods aim to produce a valid,
explainable result according to the solver's rules. Unlike generative AI,
they do not create an answer by predicting likely text or images. They are
useful when a decision must satisfy precise requirements, such as allocating
resources, scheduling work, or finding a route, and when compliance with those requirements must be guaranteed.

## What this repository will contain

Samples are planned for multiple kinds of deterministic models and multiple
languages. Each sample should explain the problem, show how its model is
defined, and include runnable code where practical.

The first set of examples demonstrates linear and integer optimization with
Google OR-Tools:

- [Linear Programming demos](LPDemos/README.md), including [Pete's Pottery
  Paradise](LPDemos/PetesPotteryParadise/README.md), a production-planning
  problem implemented in Go and JavaScript.
