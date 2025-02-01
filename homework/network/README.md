## Go Advanced Bootcamp Week 9 Assignment

## Questions

1. Summarise several ways to decode socket sticky packets: fix length/delimiter based/length field based frame decoder. try to give an example of its application; ([Answer](#Summarise ways to decode socket sticky packets))

2. implement a decoder that decodes the goim protocol from a socket connection. ([Programme Description](#Decoder description for the goim protocol))

## Summarise the decoding of socket sticky packets

>>>tcp/udp is to transfer data by byte stream, so the protocol itself does not have the concept of frame, you need to define the frame in the application layer as well as codecs, and because the devices in the network support the maximum size of each physical packet is different, so the data being transferred will be cut or merged in the process of transferring, the packet read each time when it reaches the receiving end of the application layer may not correspond to the size of the packet when it was sent. The packets read by the application layer at the receiving end do not necessarily correspond to the packet size at the time of sending, and some of the packets may still be in transit.

### fix length

>>>Fixed-length packets, each packet has a fixed size, mainly used for single type of packets.

- Advantages: simple packing and unpacking algorithm

- Disadvantages:

    1. packet contents are often wasted if they are smaller than the fixed length.
    2. the length of the packet is not flexible and needs to be adjusted by the server and the client at the same time.

- encode: usually use left-aligned and right-complemented 0 to fill the packet.

- decode: read a fixed size, once did not read enough size of the content, continue to read until the cumulative length of the content is equal to the fixed size of a package

- Figure

! [fixLength](. /img/fixLength.png)

### delimiter based

>>> Delimiter based packets, each packet is terminated with a special symbol.

- Advantages: basically transfers the actual size

- Disadvantages:

    1. the selected special symbol cannot appear in the Body, otherwise it will cause the unpacking and sealing to be different; 2. when transferring the binary content, it is not possible to use the special symbol as the terminator.
    2. the problem of reason 1 may occur when transmitting binary content.

- encode: add the packet separator after the content to be sent.

- decode: keep reading until the separator is a packet.

- Figure

    1. `_$` is the packet separator.

! [delimiterBased](. /img/delimiterBased.png)

### length field based frame decoder

>>> Variable-length packets, fixed bits + variable-length bits, fixed bits: usually contain packet identifier, packet length, variable-length content to be passed.

- Advantages: flexible, no limitation of packet grouping, suitable for passing binary packet content

- Disadvantages:

    1. fixed bits are added to each packet, increasing the amount of content to be passed.

- encode: packet identifier (1 Byte) + length of content to be passed (2 Bytes) + content to be passed

- decode.

    1. first read the fixed length 3byte = package identifier (1 Byte) + the length of the passed content (2 Byte); 2. take byte[1 Byte], and then read the byte[1 Byte].
    2. Take byte[0] and compare it to see if it is a known identifier. 3;
    3. take byte[1:3] and convert it to numbers to get the length of the delivered content.
    4. read the content according to the length of the transfer content, one did not read a sufficient size of the content, continue to read until the cumulative length of the content read is equal to the length of the transfer content

- Diagram

    1. `0xCE` is the packet identifier, the length is fixed to 1Byte.
    2. `0x000C` is the length of the packet body, the length is fixed to 2Byte.

! [lengthFieldBasedFrame](. /img/lengthFieldBasedFrame.png)

## goim protocol decoder description

### The goim protocol diagram

! [goim](. /img/GoIM.png)

>>> The goim protocol is obtained according to [goim protocol code](https://github.com/Terry-Mao/goim/blob/e742c99ad7/api/protocol/protocol.go)

### Code description

1. `protocol.go`.
    - Protocol definition, and protocol coding and decoding
    - Read a packet of the specified size with the help of `io.ReadFull` (this method will always read to the specified size before returning, unless it encounters a read error or EOF)
    - First read the byte stream of Head size, decode Head, then calculate the Body size according to PackLen and HeadLen, and read the bytes of the Body size.

2. `server.go`
    - Start the web service
    - Write the read goim packet back to the sender.

### Instructions for use

``sh
go build
network ‘:8686’
``

### Testing

>>> Verify using goim's benchmark/client

```sh
go build benchmarks/client
client 1 1 ‘:8686’
```

>>>The following screenshot shows the test validation

! [goim protocol decode](. /img/protocol.png)
