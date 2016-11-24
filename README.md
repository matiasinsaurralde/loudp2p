# loudp2p

loudp2p is an experimental peer to peer network for music streaming. The idea comes after a few personal experiments with browser-based Bittorrent streaming, see [ng-chunked-audio](https://github.com/matiasinsaurralde/ng-chunked-audio). The rest was mostly based on [torrent-stream](https://github.com/mafintosh/torrent-stream).

## Protocol design

My current plan is to use [Protocol Buffers](https://developers.google.com/protocol-buffers/) to describe the network messages. This makes it easier to generate bindings for other languages, enabling the community to write their own clients in a standard way.

These messages will be exchanged using [gRPC](http://www.grpc.io/) over [HTTP2](https://en.wikipedia.org/wiki/HTTP/2). These will involve standard data structures for peer discovery (announcements), search (index lookup) and stream initiation (*Peer A would like to stream file X from Peer B and Peer C, etc.*).

The data streams will use dedicated UDP sockets, on random ports, just like other P2P applications do. If UDP sockets aren't available, an alternative approach could wrap the data chunks as Protocol Buffer messages and send them over gRPC/HTTP2. Think about browser streaming.

A network peer will have a hybrid role most of time, e.g participating as a server and a client at the same time. See the section about peer roles.

## Bootstrapping

An initial list of peers will be used for bootstrapping. Peer discovery strategies need more investigation.

Peer discovery should be implemented in a way that guarantees the unstructured operation of the network.

## Peer roles

A peer that participates as a "server" will reply general queries and stream requests if the data is available locally. It should also receive and handle announcement messages from other peers, and keep a track of them during its lifetime.

## Light clients

In the blockchain world, a light client represents a network client that doesn't need a copy of the whole blockchain to participate. In the context of this project, think about a mobile phone with a 3G connection, running a minimal client with support for peer discovery and streams.

## Licensing

[MIT](LICENSE.md)
