---
description: Definition of StreamingFast Substreams
---

# Definition

### What is Substreams?

Substreams is an exceptionally powerful processing engine that consumes streams of rich blockchain data that can be refined and shaped for painless digestion by end-user applications, such as decentralized exchanges.

Substreams can be used to populate any kind of data store and also employs extremely powerful parallelization techniques to process huge, and ever-growing, blockchain histories.

Substreams can be scaled horizontally resulting in a massive reduction of processing time, and ultimately cost, through the addition of multiple machines.

Communities can combine Substreams data refinement strategies to form compounding levels of data richness and availability.

Substreams brings a handful of new concepts to the wider ecosystem surrounding The Graph and its subgraphs.&#x20;

Subgraphs require one or more smart contract addresses to collect data to query. Substreams don’t require _ANY_ contract addresses and the _**entire chain** with history_ is available to query.

Substreams was inspired by traditional large-scale data systems now _fused_ with the novelties of blockchain.

Substreams are defined in modules written in the Rust programming language and utilize Google Protocol Buffer technology.

More detailed aspects of what Substreams is, and can do, in contrast to what it's not, and can't do, is helpful to gain an even deeper context of the product.

#### _Substreams **is:**_

* a streaming-first system based on gRPC, protobuf, and StreamingFast Firehose,
* a highly cacheable and parallelizable remote code execution framework,&#x20;
* composable down to individual modules,
* enables the community to build higher-order modules with great ease,
* being fed by deterministic blockchain data and is therefor deterministic.

#### _Substreams is **NOT:**_

* a relational database,
* REST service,
* concerned directly with how data is stored,
* a general-purpose _non-deterministic_ event stream processor.

The _word_ Substreams refers to:

* a wink to Subgraphs,
* a plurality of _streams_, each in the form of a _module,_
* packed in a single package, but streamable individually a _sub_unit of a package,
* _streams_ composed from imported modules, blended, enriched or refined together (as in _sub_ or downstream component),
* a manifest or package will usually contain more than one module, and/or import one or more modules. It is therefore fitting to talk about a package being a _Substreams_ package.

The Substreams _engine_ is completely agnostic of underlying blockchain protocols and works solely on data extracted from nodes using the Firehose.&#x20;

Different protocols have different chain-specific extensions, such as Ethereum, which expose `eth_calls`.

**Substreams in More Detail**&#x20;

Substreams enables blockchain developers to write Rust modules that compose data streams alongside the community. The end result of community-developed solutions provides far more meaningful blockchain data than ever before.

Substreams provides extremely high-performance indexing by virtue of parallelization, in a streaming-first fashion. These powerful parallelization techniques enable efficient processing of enormous blockchain histories.

Substreams is horizontally scalable presenting the opportunity to reduce processing time simply by adding more computing power, or machines.

Substreams has all the benefits of Firehose, like low-cost caching and archiving of blockchain data, high throughput processing, and cursor-based reorgs handling.

Substreams is the successor of [StreamingFast Sparkle](https://github.com/streamingfast/sparkle). The current Substreams iteration enables greater composability, provides similar powers of parallelization. Basically, Substreams is a _much_ simpler model to work with.
