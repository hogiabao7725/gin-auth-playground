# go-ticket-engine

A REST API authentication service built with Go and Gin.

## Learning Reference

This project attempts to follow clean architecture principles but the implementation is **incomplete and has design inconsistencies**.

It works functionally, but the code organization is not optimal for extensibility and maintainability. It serves as a learning example of what not to do in certain areas.

## ✅ Proper Clean Architecture Implementation

For a **complete and correct** clean architecture reference, use:

- [go-auth-playground](https://github.com/hogiabao7725/go-auth-playground)

This is the recommended version with proper layering, clear separation of concerns, and extensible design.

## What's in This Project

This repo uses:

- Gin web framework
- PostgreSQL + sqlc
- Zerolog logging
- JWT + Bcrypt
- Role-based access control

Use it as a reference for the technologies, but prefer `go-auth-playground` for architectural patterns.
