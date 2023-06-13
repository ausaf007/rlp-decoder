<h1 align="center">RLP Decoder</h1>

<h3 align="center">Command Line Tool To Decode RLP Encoded String </h3>

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ul>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#tech-stack">Tech Stack</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#how-to-use">How to use?</a></li>
  </ul>
</details>

## About The Project

This is a simple command line application that recursively decodes a given RLP encoded string. [Learn more](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/) about RLP Encoding/Decoding.  

## Tech Stack

[![](https://img.shields.io/badge/Built_with-Go-green?style=for-the-badge&logo=Go)](https://go.dev/)

## Prerequisites

Download and install [Golang 1.19](https://go.dev/doc/install) (or higher).  

## How To Use?

1. Navigate to `rlp-decoder/`:
   ``` 
   cd /path/to/folder/rlp-decoder/
   ```
2. Get dependencies:
   ``` 
   go mod tidy
   ```
3. Run the app:
   ``` 
   go run main.go 
   # use "--verbose" flag to get additional logs
   go run main.go --verbose 
   ```
4. CD into `rlp/` to run tests: 
   ``` 
   cd rlp/
   go test
   ```

Thank you!