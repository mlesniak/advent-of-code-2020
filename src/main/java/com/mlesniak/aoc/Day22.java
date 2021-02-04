package com.mlesniak.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;

public class Day22 {
    public static void main(String[] args) throws Exception {
        var deck1 = new LinkedList<Integer>();
        var deck2 = new LinkedList<Integer>();
        parseDeck(deck1, deck2);

        playGame(1, 1, deck1, deck2);

        // Compute score
        printScore(deck1, deck2);
    }

    private static void parseDeck(LinkedList<Integer> deck1, LinkedList<Integer> deck2) throws IOException {
        var lines = Files.readAllLines(Path.of("input/22.txt"));
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
    }

    private static int playGame(int game, int round, LinkedList<Integer> deck1, LinkedList<Integer> deck2) {
        var roundCache = new HashSet<String>();

//        System.out.printf("=== Game %d ===\n", game);

        while (deck1.size() != 0 && deck2.size() != 0) {
//            System.out.printf("--- Round %d (Game %d) ---\n", round, game);

            // Update cache.
            var cacheString = deck1.toString() + "|" + deck2.toString();
            if (roundCache.contains(cacheString)) {
//                System.out.println("Game already played. Player 1 won.");
                return 1;
            }
            roundCache.add(cacheString);

//            System.out.printf("Player 1 deck: %s\n", deck1);
//            System.out.printf("Player 2 deck: %s\n", deck2);

            // Otherwise, this round's cards must be in a new configuration; the players begin the
            // round by each drawing the top card of their deck as normal.
            var c1 = deck1.pop();
            var c2 = deck2.pop();

            int result;
            if (c1 <= deck1.size() && c2 <= deck2.size()) {
                // If both players have at least as many cards remaining in their deck as the value of
                // the card they just drew, the winner of the round is determined by playing a new game of Recursive Combat (see below).
                var copy1 = deck1.subList(0, c1);
                var copy2 = deck2.subList(0, c2);
//                System.out.println("Playing a sub-game to determine the winner...\n");
                result = playGame(game + 1, 1, new LinkedList<>(copy1), new LinkedList<>(copy2));
            } else {
                // Otherwise, at least one player must not have enough cards left in their deck to
                // recurse; the winner of the round is the player with the higher-value card.
//                System.out.printf("Player 1 plays: %d\n", c1);
//                System.out.printf("Player 2 plays: %d\n", c2);
                if (c1 > c2) {
                    result = 1;
                } else {
                    result = 2;
                }
            }

            if (result == 1) {
//                System.out.printf("Player 1 wins round %d of game %d!\n", round, game);
                deck1.add(c1);
                deck1.add(c2);
            } else {
//                System.out.printf("Player 2 wins round %d of game %d!\n", round, game);
                deck2.add(c2);
                deck2.add(c1);
            }
        }

        if (deck1.size() != 0) {
            return 1;
        } else {
            return 2;
        }
    }

    private static void printScore(LinkedList<Integer> deck1, LinkedList<Integer> deck2) {
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
