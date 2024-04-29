# Forge - Framework to write ETL Pipelines driven by a central config store.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Design Patterns](#design-patterns)
- [Todos and Future Plans](#todos-and-future-plans)

## Overview
This project is aimed at providing an opiniated way of writing ETL Pipelines in a config driven way. The code follows a few design patterns in order to make the code lean and easy to write. There is an accompanying CLI which generates boilerplate code and constructs an option factory.

## Features
 1. Framework and barebone boilerplate to write any ETL code.  
 2. Provides a neat way to use a central config store to compose ETL pipelines using just configuration.

## Design Patterns
1. Uses a Chain of Responsibility design pattern to execute each step of the ETL Pipeline.
2. Uses a Factory Pattern to use the central config store to compose ETL pipelines in different ways.







