import fs from "fs";

function main() {
    const input = fs.readFileSync("inputs/day01.txt", "utf8");
    const lines = input.split("\n");
    const left = [];
    const right = [];
    let sum = 0;

    for (let line of lines) {
        const [a, b] = line.split("   ").map(Number);
        left.push(a);
        right.push(b);
    }

    left.sort((a, b) => a - b);
    right.sort((a, b) => a - b);

    for (let i = 0; i < left.length; i++) {
        sum += left[i] > right[i] ? left[i] - right[i] : right[i] - left[i];
    }

    console.log(sum); // 1222801
}

main();
