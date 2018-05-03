# yaaibig

[![Maintainability](https://api.codeclimate.com/v1/badges/b463e3ee73b8b3c65a41/maintainability)](https://codeclimate.com/github/Marahin/yaaibig/maintainability)

### Yet Another "Assembly" Interpreter, But In Go

## What this is

This is an interpreter for a very basic, simplified "assembly" dialect. **It does not comply with any standards, although it bases on some of the most popular instructions**.  
**This was made purely for fun.**

## What this is NOT

* a _real_ Assembly,  
* _efficient_ way to write code,
* a compiler,
* production ready.

## Purpose(s) 

I wanted to see how an interpreter for a simple language can be done in Go, assuming not so sophisticated instructions. Assembly was a perfect choice as it follows a very straightforward syntax and instructions. 

What I also love about assembly is how primitive the code has to be, due to the limitations. This way it was also so easy to implement & test the instructions, as - for instance - the Fibonacci or Factorial program consists of three most common used operators (`mov`, `mul`, `add`, `jnz`). 

Also thanks to this being in Golang, it can easily be used for various other purposes, such as:  

* deploying to Lambda, making a lightning talk about it <*wink wink*>,  
* explaining how the basics of Assembly work; sample programs

## Language definition

### Definitions

Following descriptions are purely for the convenience of having same understanding of yaaibig's internals.

#### INSTRUCTION

INSTRUCTION is a single line of code, exactly how the programmer perceives it. It can contain code, it can be empty, or it can contain comment. 

Current INSTRUCTION is stored in instruction REGISTER (see REGISTRY section). 

#### REGISTRY

REGISTRY is a cell which can contain VALUE. Registers are identified by a single character (A-Z). Registers identified by lower case characters _should_ be considered as _internal_, that means those can be used for interpreter's purposes.    
Currently, only following registers are reserved:  

* `i` - INSTRUCTION registry, which contains current INSTRUCTION index,  
* `m` - memory registry, to which some operands write

Above registers **SHOULD** be considered as internal, but **it it not enforced**. User is still **able to read, write and change values of internal registers**.  
For example, changing the INSTRUCTION registry value will change which INSTRUCTION will be executed next.

#### VALUE

**In current implementation** VALUE is an `integer`.

### Operands

### `mov`

`mov` takes two arguments (`ARG1`, `ARG2`).  
`ARG1` is a REGISTRY, and `ARG2` can be either REGISTRY or VALUE.  
If `ARG2` is a REGISTRY, then value of REGISTRY `ARG1` will be set to value of `ARG2` REGISTRY.
If `ARG2` is a VALUE, then value of REGISTRY `ARG1` will be set to `ARG2`.  

### `mul`

`mul` takes two arguments (`ARG1`, `ARG2`). It has a common implementation like `imul` in the real assembly dialects.  
For registers, it will evaluate `ARG1` and `ARG2`, as values stored in given registry.  
Then it executes `ARG1*ARG2` and stores the result in memory REGISTRY.

### `add`

`add` takes two arguments (`ARG1`, `ARG2`).  
For registers, it will evaluate `ARG1` and `ARG2`, as values stored in given registry.  
Then it executes `ARG1+ARG2` and stores the result in memory REGISTRY.

### `jnz`

`jnz` takes one argument (`ARG`).  
`ARG` has to be an integer value. If memory REGISTRY (`m`) is non-zero, then INSTRUCTION REGISTRY will change to `ARG` (effectively changing the next instruction to be executed).

### `jmp`

**NOT IMPLEMENTED YET**
