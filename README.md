# [![Sia Logo](https://raw.githubusercontent.com/veritascoin/Sia/master/veritascoin.png)](http://sia.tech) v1.3.2 (Capricorn)

[![Build Status](https://travis-ci.org/NebulousLabs/Sia.svg?branch=master)](https://travis-ci.org/NebulousLabs/Sia)
[![GoDoc](https://godoc.org/github.com/NebulousLabs/Sia?status.svg)](https://godoc.org/github.com/NebulousLabs/Sia)
[![Go Report Card](https://goreportcard.com/badge/github.com/NebulousLabs/Sia)](https://goreportcard.com/report/github.com/NebulousLabs/Sia)

VeritasCoin is a new decentralized cloud storage platform that radically alters the
landscape of cloud storage for Law firms. By leveraging smart contracts, client-side
encryption, and sophisticated redundancy (via Reed-Solomon codes), VeritasCoin allows
users to safely store their data with hosts that they do not know or trust.
The result is a cloud storage marketplace where hosts compete to offer the
best service at the lowest price. And since there is no barrier to entry for
hosts, anyone with spare storage capacity can join the network and start
making money.

![UI](https://raw.githubusercontent.com/veritascoin/Sia/master/veritascoin.png)

Traditional cloud storage has a number of shortcomings. Users are limited to a
few big-name offerings: Google, Microsoft, Amazon. These companies have little
incentive to encrypt your data or make it easy to switch services later. Their
code is closed-source, and they can lock you out of your account at any time.

We believe that users should own their data. VeritasCoin achieves this by replacing
the traditional monolithic cloud storage provider with a blockchain and a
swarm of hosts, each of which stores an encrypted fragment of your data. Since
the fragments are redundant, no single host can hold your data hostage: if
they jack up their price or go offline, you can simply download from a
different host. In other words, trust is removed from the equation, and
switching to a different host is painless. Stripped of these unfair
advantages, hosts must compete solely on the quality and price of the storage
they provide.

VeritasCoin can serve as a replacement for personal backups, bulk archiving, content
distribution, and more. For developers, VeritasCoin is a low-cost alternative to
Amazon S3. Storage on VeritasCoin is a full order of magnitude cheaper than on S3,
with comparable bandwidth, latency, and durability. VeritasCoin works best for static
content, especially media like videos, music, and photos.

Distributing data across many hosts automatically confers several advantages.
The most obvious is that, just like BitTorrent, uploads and downloads are
highly parallel. Given enough hosts, VeritasCoin can saturate your bandwidth. Another
advantage is that your data is spread across a wide geographic area, reducing
latency and safeguarding your data against a range of attacks.

It is important to note that users have full control over which hosts they
use. You can tailor your host set for minimum latency, lowest price, widest
geographic coverage, or even a strict whitelist of IP addresses or public
keys.

At the core of VeritasCoin is a blockchain that closely resembles Bitcoin.
Transactions are conducted in VeritasCoin, a cryptocurrency. The blockchain is
what allows VeritasCoin to enforce its smart contracts without relying on centralized
authority.

