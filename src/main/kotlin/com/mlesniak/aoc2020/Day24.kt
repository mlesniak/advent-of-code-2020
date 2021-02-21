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

enum class Color {
    Black,
    White
}

// We're using https://www.redblobgames.com/grids/hexagons/#coordinates with offset rows.
class Coordinate(val x: Int, val y: Int) {
    var color: Color = Color.White

    override fun toString(): String {
        return "($x, $y) / $color"
//        return "($x, $y)"
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as Coordinate

        if (x != other.x) return false
        if (y != other.y) return false

        return true
    }

    override fun hashCode(): Int {
        var result = x
        result = 31 * result + y
        return result
    }

    fun copy(): Coordinate {
        val c = Coordinate(x, y)
        c.color = color
        return c
    }
}

class Day24 {
    fun main() {
        var tiles = mutableMapOf<Coordinate, Coordinate>()
        val directions = parse()
        prepareInitialCells(directions, tiles)

        for (i in 0..100) {
            val numBlack = tiles.keys.filter { it.color == Color.Black }.size
            println("Day $i: $numBlack")
            tiles = step(tiles)
        }
    }

    private fun prepareInitialCells(directions: List<Directions>, tiles: MutableMap<Coordinate, Coordinate>) {
        directions.forEachIndexed() { row, directions ->
            val endTile = run(directions)
            val t = tiles.get(endTile)
            if (t != null) {
                if (t.color == Color.White) {
                    t.color = Color.Black
                } else {
                    t.color = Color.White
                }
            } else {
                endTile.color = Color.Black
                tiles[endTile] = endTile
            }
        }
    }

    private fun step(tiles: MutableMap<Coordinate, Coordinate>): MutableMap<Coordinate, Coordinate> {
        val newTiles = mutableMapOf<Coordinate, Coordinate>()

        // For all existing (blacks and white) tiles.
        for (tile in tiles.keys) {
            val neighbors = countBlackNeighbors(tiles, tile)
            val nt = tile.copy()
            when (nt.color) {
                Color.Black ->
                    if (neighbors == 0 || neighbors > 2) {
                        nt.color = Color.White
                    }
                Color.White ->
                    if (neighbors == 2) {
                        nt.color = Color.Black
                    }
            }
            newTiles[nt] = nt
        }

        // For all plates which are around any plate.
        for (cur in tiles.keys) {
            // Check all surrounding tiles.
            val surroundingTiles = getNeighbors(cur)
            // For each of the neighbors, check if it should be switched.
            for (neighbor in surroundingTiles) {
                val numNeighbors = countBlackNeighbors(tiles, neighbor)
                val nt = neighbor.copy()
                when (nt.color) {
                    Color.Black ->
                        if (numNeighbors == 0 || numNeighbors > 2) {
                            nt.color = Color.White
                            newTiles[nt] = nt
                        }
                    Color.White ->
                        if (numNeighbors == 2) {
                            nt.color = Color.Black
                            newTiles[nt] = nt
                        }
                }
            }
        }

        return newTiles
    }

    private fun run(d: Directions): Coordinate {
        var cur = Coordinate(0, 0)

//        println("\n\n")
        d.forEach { dir ->
            val old = cur
            cur = when (dir) {
                Direction.East -> Coordinate(cur.x + 1, cur.y)
                Direction.West -> Coordinate(cur.x - 1, cur.y)

                Direction.NorthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x - 1, cur.y - 1)
                Direction.NorthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x + 1 else cur.x, cur.y - 1)

                Direction.SouthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x - 1, cur.y + 1)
                Direction.SouthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x + 1 else cur.x, cur.y + 1)
            }
//            println("$old + $dir = $cur")
        }

        return cur
    }

    private fun countBlackNeighbors(tiles: MutableMap<Coordinate, Coordinate>, cur: Coordinate): Int {
        var counter = 0
        val ns = listOf(
            Direction.East,
            Direction.SouthEast,
            Direction.SouthWest,
            Direction.West,
            Direction.NorthWest,
            Direction.NorthEast
        )

        for (d in ns) {
            val n = computeCoordinate(cur, d)
            val t = tiles.get(n)
            if (t != null) {
                if (t.color == Color.Black) {
                    counter++
                }
            }
        }

        return counter
    }

    private fun getNeighbors(cur: Coordinate): Set<Coordinate> {
        val tiles = mutableSetOf<Coordinate>()

        var counter = 0
        val ns = listOf(
            Direction.East,
            Direction.SouthEast,
            Direction.SouthWest,
            Direction.West,
            Direction.NorthWest,
            Direction.NorthEast
        )

        for (d in ns) {
            val n = computeCoordinate(cur, d)
            tiles.add(n)
        }

        return tiles
    }


    private fun computeCoordinate(cur: Coordinate, dest: Direction): Coordinate {
        return when (dest) {
            Direction.East -> Coordinate(cur.x + 1, cur.y)
            Direction.West -> Coordinate(cur.x - 1, cur.y)

            Direction.NorthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x - 1, cur.y - 1)
            Direction.NorthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x + 1 else cur.x, cur.y - 1)

            Direction.SouthWest -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x else cur.x - 1, cur.y + 1)
            Direction.SouthEast -> Coordinate(if ((cur.y % 2).absoluteValue == 1) cur.x + 1 else cur.x, cur.y + 1)
        }
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