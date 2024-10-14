const BoardSize = 8;

enum Color {
  White = "White",
  Black = "Black",
}

interface Game {
  board: (Piece | null)[][];
  activeColor: Color;
}

interface Piece {
  type: string;
  color: Color;
  symbol: string;
}

interface MoveRequest {
  fromX: number;
  fromY: number;
  toX: number;
  toY: number;
  color: Color;
}

async function makeMove(move: MoveRequest): Promise<string> {
  const response = await fetch("/move", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(move),
  });

  if (!response.ok) {
    throw new Error("Move failed");
  }

  return response.json();
}
