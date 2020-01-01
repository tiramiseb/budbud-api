# Bud(get)-Bud(dy) API

Bud-Bud is a (very) opiniated personal finance manager, running as a web service.

This is the API.

## Opinions

* No automatic operations fetching, because if you want to track your budget, you have to take a look
* Repetition logic is monthly, nothing else
* However, you may want to repeat some expenses only on some months of the year (for instance, some insurance you would pay only once a year)
* The only income is your salary or such stuff: medical reimbursements (for instance) are not accounted as income, but as counterparts of other operations (or, in other words, be related to one or multiple other operations)
* You may relate an operation to one or more other operations (typically, medical reimbursements that relate to multiple doctor appointments)
* Categories are split into supercategories, but there is no other hierarchy

Personal budget management is not accounting, and sometimes you may have weird or "useless" operations on your account, that's why there is no obligation to relate exactly to your real account amounts, and you can ignore some operations in your budget.

## In-development warning

I have hard-coded a very permissive CORS. It will probably be switchable with an option. Or completely removed for production. Something like that...

Well, many things must be read from a config file instead of being hard-coded...

Quick&dirty:

```sh
cd cmd/budbud
go run main.go
```