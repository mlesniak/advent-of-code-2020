package com.mlesniak.aoc2020

import java.nio.file.Files
import java.nio.file.Path
import java.util.*
import kotlin.math.absoluteValue

enum class Direction(val tag: String) {
    East("e"),
    SouthEast("se"),
    SouthWest("sw"),
    West("w"),
    NorthWest("nw"),
    NorthEast("ne"),
}

typealias Directions = List<Direction>

data class Coordinate(val x: Int, val y: Int) {
    override fun toString(): String {
        return "($x, $y)"
    }
}

class Day24 {
    fun main() {
        val directions = parse()

        // We're using https://www.redblobgames.com/grids/hexagons/#coordinates with offset rows.
        val tiles = mutableSetOf<Coordinate>()

        directions.forEachIndexed() { row, directions ->
            val endTile = run(directions)
            if (tiles.contains(endTile)) {
                println("Back to white    : $row $endTile ")
                tiles.remove(endTile)
            } else {
                println("Flipping to black: $row $endTile ")
                tiles.add(endTile)
            }
        }

        println("Black tiles = ${tiles.size}")
    }

    private fun run(d: Directions): Coordinate {
        var cur = Coordinate(0, 0)

        println("\n\n")
        d.forEach { dir ->
            val old = cur
            cur = when (dir) {
                Direction.East -> Coordinate(cur.x + 1, cur.y)
                Direction.West -> Coordinate(cur.x - 1, cur.y)

                Direction.NorthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x-1, cur.y-1)
                Direction.NorthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x+1 else cur.x, cur.y-1)

                Direction.SouthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x-1, cur.y+1)
                Direction.SouthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x+1 else cur.x, cur.y+1)
            }
            println("$old + $dir = $cur")
        }

        return cur
    }

    private fun parse(): List<Directions> {
        val lines = Files.readAllLines(Path.of("input/24.txt"))
        return lines.map { parseLine(it) }
    }

    private fun parseLine(line: String): Directions {
        val sb = StringBuilder(line)
        val dirs = mutableListOf<Direction>()

        while (sb.isNotEmpty()) {
            for (d in Direction.values()) {
                if (sb.startsWith(d.tag)) {
                    //println("${sb.substring(0, d.tag.length)} -> $d")
                    dirs.add(d)
                    sb.delete(0, d.tag.length)
                }
            }
        }

        return Collections.unmodifiableList(dirs)
    }
}

fun main() {
    Day24().main()
}