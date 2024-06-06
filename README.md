# Forge - Framework to write ETL Pipelines driven by a central config store.

## Table of Contents
- [Forge - Framework to write ETL Pipelines driven by a central config store.](#forge---framework-to-write-etl-pipelines-driven-by-a-central-config-store)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Demo](#demo)
  - [Design Patterns](#design-patterns)
  - [Todos](#todos)

## Overview
This project is aimed at providing a framework/structure for writing ETL Pipelines driven by a central config store. The code follows a few design patterns to make the code lean and easy to write. There is an accompanying CLI that generates boilerplate code and constructs an option factory.

## Features
 1. Framework and barebone boilerplate to write any ETL code.  
 2. Provides a neat way to use a central config store to compose ETL pipelines using just configuration.
 3. Accompanying this repository there is a
    -   [CLI tool which does boilerplate code generation for different ETL Steps](https://github.com/harish876/forge/tree/main/cli)
    -   [An LSP that provides code completion and goto definition features for the configs belonging to a specific job](https://github.com/harish876/forge-lsp)

## Demo
<video width="600" controls>
  <source src="https://harish876.github.io/forge/cli_demo.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

If the video does not play, [click here to watch the CLI Demo](https://harish876.github.io/forge/cli_demo.mp4).

<video width="600" controls>
  <source src="https://harish876.github.io/forge/lsp_demo.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

If the video does not play, [click here to watch the LSP Demo](https://harish876.github.io/forge/lsp_demo.mp4).

## Design Patterns
1. Uses a Chain of Responsibility design pattern to execute each step of the ETL Pipeline.
2. The idea is to create a linked list of jobs and then provide flexibility to the initiator of the linked list to execute each
   step and traverse through the list of jobs in an iterative fashion. This is particularly useful in paginating, and streaming a large
   dataset.
3. Uses a Factory Pattern to use the central config store to compose ETL pipelines in different ways.


## Todos
1. Add Video Documentation to this repository for better presentation
2. Create a Project ecosystem on Git Hub and make the CLI tool and LSP tool into individual repositories.
3. Add a utility to add merge runtime CLI args with the original configs.




