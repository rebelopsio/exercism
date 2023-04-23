const values = new Map<string, number>([
  ["black", 0],
  ["brown", 1],
  ["red", 2],
  ["orange", 3],
  ["yellow", 4],
  ["green", 5],
  ["blue", 6],
  ["violet", 7],
  ["grey", 8],
  ["white", 9],
]);

export function decodedValue(colors: string[]): number {
  return parseInt(
    colors
      .slice(0, 2)
      .map((color) => values.get(color))
      .join("")
  );
}
