trylock - TryLock implementation for Go
=======================================

[![Build Status](https://travis-ci.org/LK4D4/trylock.svg?branch=master)](https://travis-ci.org/LK4D4/trylock)
[![GoDoc](https://godoc.org/github.com/LK4D4/trylock?status.svg)](https://godoc.org/github.com/LK4D4/trylock)

Please see godoc for documentation. It uses unsafe, which is sorta "unsafe", but
should work until `sync.Mutex` will change its layout (I hope it never will).
