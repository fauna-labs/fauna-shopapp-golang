Sample shop API using Fauna and Go
=============

#### Table of Contents
* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Set up your Fauna database](#set-up-your-fauna-database)
* [Run the app locally](#run-the-app-locally)

## Overview
This is a sample Go REST API an e-commerce application that uses [Fauna](https://docs.fauna.com/) as a database, and Fauna's [Go driver](https://github.com/fauna/faunadb-go).

## Prerequisites
You need to install Go:  
https://golang.org/doc/install

Supported Go versions:
- 1.16 and higher

The next step uses the [Fauna Dashboard](https://dashboard.fauna.com) to set up your database. Alternatively, you could use the [Fauna Shell/CLI tool](https://github.com/fauna/fauna-shell).

## Set up your Fauna database

1. [Sign up for free](https://dashboard.fauna.com/accounts/register) or [log in](https://dashboard.fauna.com/accounts/login) at [dashboard.fauna.com](https://dashboard.fauna.com/accounts/register).
2. Click [CREATE DATABASE], name it "shopapp", select a region group (e.g., "Classic"), and click [CREATE].
3. Click the [SECURITY] tab at the bottom of the left sidebar, and [NEW KEY].
4. Create a Key with the default Role of "Admin" selected, and copy/paste the secret somewhere safe. It will not be displayed again, and you'll need this secret to start the app in the next section.
5. Navigate to your "shopapp" database by clicking on it from the main page of your [Fauna Dashboard](https://dashboard.fauna.com)
6. Click [New Collection], name it "categories", and click [Save].
7. Do the same to create a "products" collection.

## Run the app locally
1. To access your Fauna database, you'll need to export your secret via an environment variable. Get `your-secret-key` from the secret you copied in the [Set up your Fauna database](#set-up-your-fauna-database) section, and run the following in your terminal:
```
export FAUNA_SECRET_KEY=your-secret-key
```
2. Start a server by running:
```
go run main.go
```
3. When you start your server, a Swagger UI should be available at http://localhost:8080/swagger/index.html#/ where you can test the following REST endpoints:
    - List of categories: [GET] http://localhost:8080/categories
    - Create a new category: [POST] http://localhost:8080/categories
    - Create a new product: [POST] http://localhost:8080/product
