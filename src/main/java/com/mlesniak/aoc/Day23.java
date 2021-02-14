package com.mlesniak.aoc;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Element {
    int value;
    Element next;

    @Override
    public String toString() {
        return String.format("%d", value);
//        return "Element{" +
//                "value=" + value +
//                ", next=" + next.value +
//                '}';
    }
}

class CircleArray {
    private Element root;
    private int highestValue;
    private int lowestValue;

    private Map<Integer, Element> table;

    public CircleArray(String input, Integer max) {
        highestValue = Integer.MIN_VALUE;
        lowestValue = Integer.MAX_VALUE;

        table = new HashMap<>();

        Element prev = null;
        for (int i = 0; i < input.length(); i++) {
            var value = Integer.parseInt(String.valueOf(input.charAt(i)));
            if (value > highestValue) {
                highestValue = value;
            }
            if (value < lowestValue) {
                lowestValue = value;
            }
            if (i == 0) {
                root = new Element();
                root.value = value;
                prev = root;
                table.put(root.value, root);
                continue;
            }

            var node = new Element();
            node.value = value;
            prev.next = node;
            prev = node;

            table.put(node.value, node);
        }

        if (max != null) {
            for (int i = highestValue+1; i <= max; i++) {
                var node = new Element();
                node.value = i;
                prev.next = node;
                prev = node;
                table.put(node.value, node);
            }
        }

        prev.next = root;
    }

    public int getHighestValue() {
        return highestValue;
    }

    private void updateEdgeValues(Element current, ArrayList<Element> threes) {
        highestValue = 1_000_000;
        while (threes.stream().anyMatch(d -> d.value == highestValue)) {
            highestValue--;
        }

        lowestValue = 1;
        while (threes.stream().anyMatch(d -> d.value == lowestValue)) {
            lowestValue++;
        }
    }

    public String result() {
        var sb = new StringBuilder();
        var current = root;
        while (current.value != 1) {
            current = current.next;
        }
        current = current.next;

        while (current.value != 1) {
            sb.append(current.value);
            current = current.next;
        }

        return sb.toString();
    }

    public long result2() {
        var sb = new StringBuilder();
//        var current = root;
//        while (current.value != 1) {
//            current = current.next;
//        }
        var current = table.get(1);

        current = current.next;
        long v1 = current.value;
        current = current.next;
        long v2 = current.value;

        return v1 * v2;
    }

    public Element root() {
        return root;
    }

    public void insert(Element cur, List<Element> threes) {
        var tmp = cur.next;
        cur.next = threes.get(0);
        threes.get(2).next = tmp;

    }

    public List<Element> takeThree(Element cur) {
        var tmp = cur;
        var result = new ArrayList<Element>();
        for (int i = 0; i < 3; i++) {
            cur = cur.next;
            result.add(cur);
        }
        tmp.next = cur.next;
        updateEdgeValues(tmp, result);

        return result;
    }

    public Element findDestinationCup(Element current, List<Element> threes) {
        var targetValue = current.value - 1;
        var found = false;
        while (!found) {
            found = true;
            final var t = targetValue;
            if (threes.stream().anyMatch(d -> d.value == t)) {
                targetValue -= 1;
                found = false;
            }
            if (targetValue < lowestValue) {
                targetValue = highestValue;
            }
        }

        // Not efficient, but :shrug: ...
        // TODO(mlesniak) use map here
//        var tmp = root;
//        while (tmp != null) {
//            if (tmp.value == targetValue) {
//                return tmp;
//            }
//            tmp = tmp.next;
//            if (tmp == root) {
//                break;
//            }
//        }
        var element = table.get(targetValue);
        if (element == null) {
            System.out.println("break");
        }
        return element;
//        return null;
    }

    public String toString() {
        var sb = new StringBuilder();
        var current = root;
        while (current != null) {
            sb.append(String.format("%d ", current.value));
            current = current.next;
            if (current == root) {
                break;
            }
        }
        sb.deleteCharAt(sb.length() - 1);

        return sb.toString();
    }
}

/**
 * gradle build
 * java -cp build/libs/aoc-2020-1.0-SNAPSHOT.jar com.mlesniak.aoc.Day23
 */
public class Day23 {
    public static void main(String[] args) {
//        var input = "389125467";
        var input = "716892543";
        var elements = new CircleArray(input, 1_000_000);
        var cur = elements.root();

        var moves = 10_000_000;

        for (int i = 1; i <= moves; i++) {
            var now = System.currentTimeMillis();
            var threes = elements.takeThree(cur);
            var destination = elements.findDestinationCup(cur, threes);
            if (destination == null) {
                System.out.println("null");
            }
            elements.insert(destination, threes);
            cur = cur.next;
            if (i % 1000 == 0) {
                var percent = ((float) i / moves) * 100.0;
                System.out.printf("--- move %d / %g ---: %d\n", i, percent, System.currentTimeMillis() - now);
            }
        }

        System.out.println(elements.result2());
    }
}
