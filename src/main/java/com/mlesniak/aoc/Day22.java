package com.mlesniak.aoc;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.LinkedList;
import java.util.List;

public class Day22 {
    public static void main(String[] args) throws Exception {
        var lines = Files.readAllLines(Path.of("input/22.txt"));

        var deck1 = new LinkedList<Integer>();
        var deck2 = new LinkedList<Integer>();


        var deck = deck1;
        for (int i = 1; i < lines.size(); i++) {
            var line = lines.get(i);
            if (line.isEmpty()) {
                i++;
                deck = deck2;
                continue;
            }
            deck.add(Integer.parseInt(line));
        }

        while (deck1.size() != 0 && deck2.size() != 0) {
            System.out.println(deck1);
            System.out.println(deck2);
            var c1 = deck1.pop();
            var c2 = deck2.pop();
            if (c1 > c2) {
                deck1.add(c1);
                deck1.add(c2);
            } else {
                deck2.add(c2);
                deck2.add(c1);
            }
        }

        // Compute score
        int score;
        if (deck1.size() != 0) {
            score = compute(deck1);
        } else {
            score = compute(deck2);
        }
        System.out.println(score);
    }

    private static int compute(List<Integer> deck) {
        var sum = 0;
        for (int i = 0; i < deck.size(); i++) {
            sum += deck.get(i) * (deck.size()-i);
        }
        return sum;
    }
}
