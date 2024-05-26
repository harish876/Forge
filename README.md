# Forge - Framework to write ETL Pipelines driven by a central config store.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Design Patterns](#design-patterns)
- [Todos] (#Todos)

## Overview
This project is aimed at providing an opiniated way of writing ETL Pipelines in a config driven way. The code follows a few design patterns in order to make the code lean and easy to write. There is an accompanying CLI which generates boilerplate code and constructs an option factory.

## Features
 1. Framework and barebone boilerplate to write any ETL code.  
 2. Provides a neat way to use a central config store to compose ETL pipelines using just configuration.
 3. Accompanying this repository there is a
    -   CLI tool which does boilerplate code generation for different ETL Steps.
    -   An LSP which provides code completion and goto definition features for the configs belonging to a specific job.

## Design Patterns
1. Uses a Chain of Responsibility design pattern to execute each step of the ETL Pipeline.
2. The idea is to create a linked list of jobs and then provide flexibility to the initiator of the linked list to execute each
   step and traverse through the list of jobs in an iterative fashion. This is particulary useful in paginating, streaming a large
   dataset.
3. Uses a Factory Pattern to use the central config store to compose ETL pipelines in different ways.

## Todos
1. Add Video Documentation to this repository for better presentation
2. Create a Project ecosystem on github and make the CLI tool and LSP tool into individual repos.




