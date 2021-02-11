package com.mlesniak.aoc;

import java.util.ArrayList;
import java.util.List;

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

    public CircleArray(String input) {
        highestValue = Integer.MIN_VALUE;
        lowestValue = Integer.MAX_VALUE;

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
                continue;
            }

            var node = new Element();
            node.value = value;
            prev.next = node;
            prev = node;
        }
        prev.next = root;
    }

    private void updateEdgeValues(Element root) {
        highestValue = Integer.MIN_VALUE;
        lowestValue = Integer.MAX_VALUE;

        var current = root;
        while (current != null) {
            var val = current.value;
            if (val > highestValue) {
                highestValue = val;
            }
            if (val < lowestValue) {
                lowestValue = val;
            }
            current = current.next;
            if (current == root) {
                break;
            }
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
        updateEdgeValues(tmp);

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
        var tmp = root;
        while (tmp != null) {
            if (tmp.value == targetValue) {
                return tmp;
            }
            tmp = tmp.next;
            if (tmp == root) {
                break;
            }
        }

        return null;
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

public class Day23 {
    public static void main(String[] args) {
//        var input = "389125467";
        var input = "716892543";
        var elements = new CircleArray(input);
        var cur = elements.root();

        var moves = 100;
        for (int i = 1; i <= moves; i++) {
            System.out.printf("\n--- move %d ---\n", i);
            System.out.printf("cups: %s / cur = %s\n", elements, cur);
            var threes = elements.takeThree(cur);
            System.out.printf("pick up: %s\n", threes);
            var destination = elements.findDestinationCup(cur, threes);
            System.out.printf("destination: %d\n", destination.value);
            elements.insert(destination, threes);
            cur = cur.next;
        }

        System.out.println(elements.result());
    }
}
