package com.mlesniak.aoc2020

class Day25 {
    fun main() {
        val doorPublicKey = 5764801L
        val doorLoop = findLoop(doorPublicKey)
        println(doorLoop)

        val cardPublicKey = 17807724L
        val cardLoop = findLoop(cardPublicKey)
        println(cardLoop)

        val encryptionKey = transform(doorPublicKey, cardLoop)
        println(encryptionKey)
        // val t  = transform(7, 8)
        // println(t)
    }

    private fun findLoop(doorPublicKey: Long): Int {
        var size = 1
        while (true) {
            if (transform(7, size) == doorPublicKey) {
                return size
            }
            size++
        }
    }

    fun transform(subject: Long, loopSize: Int): Long {
        var value = 1L
        for (i in 1..loopSize) {
            value *= subject
            value %= 20201227
        }

        return value
    }
}

fun main() {
    Day25().main()
}
