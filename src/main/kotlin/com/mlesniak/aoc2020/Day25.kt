package com.mlesniak.aoc2020

class Day25 {
    fun main() {
        // 12578151
        // 5051300
        // val doorPublicKey = 5764801L
        val doorPublicKey = 12578151L
        val doorLoop = findTransform(7, doorPublicKey)
        println(doorLoop)

        // val cardPublicKey = 17807724L
        val cardPublicKey = 5051300L
        val cardLoop = findTransform(7, cardPublicKey)
        println(cardLoop)

        val encryptionKey = transform(doorPublicKey, cardLoop)
        println(encryptionKey)
        // val t  = transform(7, 8)
        // println(t)
    }

    private fun findLoop(doorPublicKey: Long): Long {
        var size = 1L
        while (true) {
            if (transform(7, size) == doorPublicKey) {
                return size
            }
            size++
        }
    }

    fun transform(subject: Long, loopSize: Long): Long {
        var value = 1L
        for (i in 1..loopSize) {
            value *= subject
            value %= 20201227
        }

        return value
    }

    fun findTransform(subject: Long,target: Long): Long {
        var value = 1L
        for (i in 1..Long.MAX_VALUE) {
            value *= subject
            value %= 20201227
            if (target == value) {
                return i
            }
        }

        return value
    }
}

fun main() {
    Day25().main()
}
