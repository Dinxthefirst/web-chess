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
}

type Square = {
  row: number;
  col: number;
};
