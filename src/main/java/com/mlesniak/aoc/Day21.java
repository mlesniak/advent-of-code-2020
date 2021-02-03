package com.mlesniak.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.stream.Collectors;

class Food {
    List<String> ingredients = new ArrayList<>();
    List<String> allergens = new ArrayList<>();

    static Food fromLine(String line) {
        var parts = line.split(" \\(contains ");
        var ing = parts[0].split(" ");
        var all = Arrays.stream(parts[1].substring(0, parts[1].length() - 1).split(","))
                .map(s -> s.trim())
                .toArray(String[]::new);

        var f = new Food();
        f.ingredients = List.of(ing);
        f.allergens = List.of(all);
        return f;
    }

    @Override
    public String toString() {
        return "Food{" +
                "ingredients=" + ingredients +
                ", allergens=" + allergens +
                '}';
    }
}

public class Day21 {
    public static void main(String[] args) throws IOException {
        List<Food> foods = parse();

        // Create a map from allergen to possible foods.
        var map = new HashMap<String, List<Food>>();
        for (var food : foods) {
            for (var allergen : food.allergens) {
                var list = map.getOrDefault(allergen, new ArrayList<>());
                list.add(food);
                map.put(allergen, list);
            }
        }

        // List all possible ingredients for all allergens.
        var candidates = new HashMap<String, Set<String>>();
        for (var entry : map.entrySet()) {
            List<List<String>> allIngredients = new ArrayList<>();
            Set<String> inAll = new HashSet<>();
            for (var food : entry.getValue()) {
                allIngredients.add(food.ingredients);
                inAll.addAll(food.ingredients);
            }
            for (var it = inAll.iterator(); it.hasNext(); ) {
                var elem = it.next();
                for (var coll : allIngredients) {
                    if (!coll.contains(elem)) {
                        it.remove();
                        break;
                    }
                }
            }
            candidates.put(entry.getKey(), inAll);
        }

        var solution = new HashMap<String, String>();
        while (candidates.size() > 0) {
            // Find entry with only one value.
            for (var iterator = candidates.entrySet().iterator(); iterator.hasNext(); ) {
                Map.Entry<String, Set<String>> entry = iterator.next();
                if (entry.getValue().size() == 1) {
                    var distinctIngredient = (String) (entry.getValue().toArray()[0]);

                    // Remove from all other entries.
                    for (var other : candidates.values()) {
                        if (other == entry.getValue()) {
                            continue;
                        }
                        other.remove(distinctIngredient);
                    }

                    solution.put(entry.getKey(), distinctIngredient);
                    iterator.remove();
                }
            }
        }

        Set<String> identified = new HashSet<>();
        for (var entry : solution.entrySet()) {
            System.out.printf("%10s %s\n", entry.getKey(), entry.getValue());
            identified.add(entry.getValue());
        }

        // Part 1
        var count = 0;
        for (var food : foods) {
            for (var ingredient : food.ingredients) {
                if (!identified.contains(ingredient)) {
                    System.out.println(ingredient);
                    count++;
                }
            }
        }
        System.out.printf("Part1: %d\n", count);

        // Part 2
        Set<String> ingredients = new HashSet<>();
        ingredients.addAll(solution.keySet());
        var sortedAllergens = ingredients.stream().sorted().collect(Collectors.toList());
        System.out.println("Part2:");
        for (var s : sortedAllergens) {
            System.out.printf("%s,", solution.get(s));
        }

    }

    private static List<Food> parse() throws IOException {
        var lines = Files.readAllLines(Path.of("input/21.txt"));
        List<Food> foods = lines.stream()
                .map(Food::fromLine)
                .collect(Collectors.toList());
        return foods;
    }
}
