# naive-banking-system

This repository is part of the Professional Programming Introduction.
The course intends to demonstrate what are the main differences on system design and implementations
between student solutions and company solutions.

## Functional Requirements

Functional Requirements describe the service that the banking management system must offer, they are subdivided into three access levels:
Admin Mode, Teller Mode, and Customer Mode.

- Customer
  - Sign in with login and password.
  - Update personal details.
  - Change password.
  - View balance.
  - View personal history of transactions.
  - Transfer money.
  - Withdraw.
  - Submit Cash.

- Customer
  - Sign in with login and password.
  - Change password.
  - Register new bank customers.
  - View customer information.
  - Manage customer accounts.

- Admin
  - Sign in with login and password.
  - View manager and customer details.
  - Add or update bank branch details.
  - Add or update manager details.

## High Level Design

The traditional way to implement bank management systems is with a monolithic architecture,
where different tasks are managed with a single unified unit.
