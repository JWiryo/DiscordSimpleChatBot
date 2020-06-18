# Discord Simple Go ChatBot

<img src="https://www.startinop.com/wp-content/uploads/2020/03/Capture-d%E2%80%99e%CC%81cran-2020-03-19-a%CC%80-20.10.14-1.png">

## Table of Contents

- [Introduction](#Introduction)
- [Name](#Name)
- [Installation](#Installation)
- [Team](#team)
- [FAQ](#faq)
- [Support](#support)
- [License](#license)

---

## Introduction

This project is an initiative by the author (JWiryo) to learn Golang in his free time by building a simple bot in his Discord channel to do several basic functions for the members of the guild.

---

## Name

The bot has been named **_Momonga_** with reference to the **_Overlord_** Anime by **_Kugane Murayama_**

Note: If you haven't watched it do give it a try

<img src="https://i.pinimg.com/474x/04/b9/92/04b992071d1c187db00d3a7745ef734e.jpg">

---

## Installation

### Prerequisites

- You need to have Golang (14.4 and above) installed in your machine

### Clone

- Clone / Fork this repo to your local machine using `https://github.com/JWiryo/DiscordSimpleChatBot.git`

### Setup

- Add the bot into your channel by following this document
  `https://discordpy.readthedocs.io/en/latest/discord.html`

### Running

Proceed to your working directory and build the bot

```bash
$ go build DiscordBot.go
```

Then run the bot

```bash
$ go run DiscordBot.go
```

Now you should see the bot coming online in your discord channel

<iframe src="https://drive.google.com/file/d/1HEN8455UpKFP_ApygwRA_ZWCw1vA_Ok-/preview"></iframe>

To switch off the program, simply press
**Ctrl + C**

## Commands

Currently Momonga supports several simple commands

#### Meme Command

```bash
/meme <<Post Number>>
```

With this command, Momonga will access Reddit's API and retrieve the n<sup>th</sup> post from Reddit _**/r/memes**_ page

#### Stock Command

```bash
/stock <<Stock Ticker>>
```

With this command, Momonga will retrieve the latest stock price of the US company identified by the ticker

Example:

```bash
/stock DIS
```

will return

> Current price of DIS is \$117.825

## References

Momonga utilises several 3<sup>rd</sup> party APIs which include

- https://github.com/turnage/graw -> Wrapper for Reddit's API for retrieval of memes
- https://github.com/goinvest/iexcloud -> Wrapper for IEX Cloud's API for latest US stock market data

Please do give these libraries their well deserved recognitions

## Future Plan

Current future plan for Momonga Bot includes:

1. Dockerization and Deployment to AWS
2. More stock calculation futures (e.g: NCAV)
3. Medium Article tutorial
