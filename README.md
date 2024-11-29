# easy-assert

# Badges

[![Build packages](https://github.com/Knoblauchpilze/easy-assert/actions/workflows/build-packages.yml/badge.svg)](https://github.com/Knoblauchpilze/easy-assert/actions/workflows/build-packages.yml)

[![codecov](https://codecov.io/gh/Knoblauchpilze/easy-assert/graph/badge.svg?token=DL8I09GHC1)](https://codecov.io/gh/Knoblauchpilze/easy-assert)

# Why this project?

## The problem statement

This project is a collection of utility functions to assert things in collections and objects.

When developing a rest application it is common to define repositories, services and controllers. Those types usually manipulate custom entities and DTOs which are (in Go) usually `struct`s.

It's also very common that those entities or DTOs ultimately are saved into or extracted from a database. And a common practice in databases is to have fields like `updated_at` or `created_at` to help tracking when rows were created.

In a testing environment and especially in integration tests, it can be quite cumbersome to assert that a certain item is in a collection of values returned by a repository or a service when all values have different timestamps.

## Isn't this a solved problem?

There are two very good building blocks for testing such scenario already:

- [testify](https://github.com/stretchr/testify) allows to write natural looking assertions for things, such as `assert.Equal(t, lhs, rhs)`.
- [go-cmp](https://github.com/google/go-cmp) makes equality checking easy and customizable by allowing to specify filters, ignored fields and more.

However, it seems that neither of those packages provide something to solve the following problem:

```go
func Test(t *testing.T) {
	type DummyStruct struct {
		Id uuid.UUID
		Name string
		UpdatedAt time.Time
	}

	values := []DummyStruct{
		{
			Id:        uuid.MustParse("071f623d-b191-40fa-9c29-b1cd8feac5a7"),
			Name:      "value1",
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			Id:       uuid.MustParse("a4fe848a-8965-4eb2-bdff-116d813d2824"),
			Name:     "somethingElse",
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
	}

	expected := DummyStruct{
		Id:   uuid.MustParse("071f623d-b191-40fa-9c29-b1cd8feac5a7"),
		Name: "value1",
	}

	// The following does not work
	assert.Contains(t, values, expected, Ignoring("UpdatedAt"))
}
```

## How to solve it?

This project adds a few convenience methods to deal with such situation. Namely: `ContainsIgnoringFields`. By calling `cmp.Equal` on all elements of a slice we can achieve the desired result. A code sample is provided below for the same test scenario as above.

```go
func Test(t *testing.T) {
	type DummyStruct struct {
		Id uuid.UUID
		Name string
		UpdatedAt time.Time
	}

	values := []DummyStruct{
		{
			Id:        uuid.MustParse("071f623d-b191-40fa-9c29-b1cd8feac5a7"),
			Name:      "value1",
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			Id:       uuid.MustParse("a4fe848a-8965-4eb2-bdff-116d813d2824"),
			Name:     "somethingElse",
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
	}

	expected := DummyStruct{
		Id:   uuid.MustParse("071f623d-b191-40fa-9c29-b1cd8feac5a7"),
		Name: "value1",
	}

	assert.True(t, easyassert.ContainsIgnoringFields(t, values, expected, "UpdatedAt"))
}
```

## What about time zones?

This is not the only thing that is solved by using this package. Another cumbersome situation to deal with is the following:

```go
func Test(t *testing.T) {
	user := User{
		Name: "my-user",
		CreatedAt: time.Now(),
	}

	// For brevity we omit the fetching of the createAt.
	dbCreatedAt, err := conn.ExecContext(
		ctx,
		`INSERT INTO users (name, created_at)
			VALUES ($1, $2)
			RETURNING created_at`,
		user.Name,
		user.CreatedAt
	)

	assert.Equal(t, user.CreatedAt, dbCreatedAt)
}
```

Depending on the settings of your database it might be that both times are not equal: this is because one might be expressed in UTC while the other one is most likely expressed in your local time zone.

Using `EqualsIgnoringFields` allows to leverage the `cmp.Equal` logic which uses the `time.Time` `Equals` method for the comparison: this correctly returns that both times are equal even though expressed in different time zones.

You can find a list of the assertions provided by this package in the [All the assertions](#all-the-assertions) section.

# Installation

To install this package and use it in our projects, just run:

```bash
go get -u github.com/Knoblauchpilze/easy-assert
```

You can then import the package in your go files as shown below:

```go
package dummy

import "github.com/KnoblauchPilze/easy-assert/assert"

func Test(t *testing.T) {
	var t1, t2 time.Time
	equal := assert.EqualsIgnoringFields(t, t1, t2)
	/* ... */
}
```

# Notes

This project was built using Go `1.23.2`. You can install from the official [download page](https://go.dev/doc/install).

As it uses generics so it can't be used with versions lower than [1.18](https://go.dev/blog/go1.18#generics).

You can clone the repository locally with:

```bash
git clone git@github.com:Knoblauchpilze/easy-assert.git`
```

# All the assertions

```go
func EqualsIgnoringFields[T any](actual T, expected T, ignoredFields ...string) bool { /* ...*/ }
```

**Brief:** returns true if `actual` and `expected` are equal (using [go-cmp](https://github.com/google/go-cmp) as a comparer) excluding the fields named in `ignoredFields`.

```go
func ContainsIgnoringFields[T any](actual []T, expected T, ignoredFields ...string) bool { /* ... */ }
```

**Brief:** returns true if `actual` contains at least one instance of `expected` using `EqualsIgnoringFields` as a comparator.

```go
func AreTimeCloserThan(t1 time.Time, t2 time.Time, distance time.Duration) bool { /* ... */ }
```

**Brief:** returns true if `t1` and `t2` are close than the distance.
