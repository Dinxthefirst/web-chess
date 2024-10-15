// import * as Types from "./types";
// TYPES
type Color = number;

enum PieceType {
  None,
  King,
  Pawn,
  Knight,
  Bishop,
  Rook,
  Queen,
  White = 8,
  Black = 16,
}

interface Piece {
  type: number;
}

interface Game {
  board: Piece[][];
  activeColor: Color;
}

interface Move {
  fromRow: number;
  fromCol: number;
  toRow: number;
  toCol: number;
}

type Square = {
  row: number;
  col: number;
};

const PieceImages: { [key in PieceType]?: { [key in Color]: string } } = {
  [PieceType.Pawn]: {
    [PieceType.White]: "/static/images/white_pawn.svg",
    [PieceType.Black]: "/static/images/black_pawn.svg",
  },
  [PieceType.Rook]: {
    [PieceType.White]: "/static/images/white_rook.svg",
    [PieceType.Black]: "/static/images/black_rook.svg",
  },
  [PieceType.Knight]: {
    [PieceType.White]: "/static/images/white_knight.svg",
    [PieceType.Black]: "/static/images/black_knight.svg",
  },
  [PieceType.Bishop]: {
    [PieceType.White]: "/static/images/white_bishop.svg",
    [PieceType.Black]: "/static/images/black_bishop.svg",
  },
  [PieceType.Queen]: {
    [PieceType.White]: "/static/images/white_queen.svg",
    [PieceType.Black]: "/static/images/black_queen.svg",
  },
  [PieceType.King]: {
    [PieceType.White]: "/static/images/white_king.svg",
    [PieceType.Black]: "/static/images/black_king.svg",
  },
};

///// EVERYHING ELSE

document.addEventListener("DOMContentLoaded", () => {
  setupNewGameButton();
});

function setupNewGameButton() {
  const newGameButton = document.getElementById("new-game-button")!;
  newGameButton.addEventListener("click", startNewGame);
}

async function startNewGame() {
  console.log("STARTING NEW GAME");

  try {
    const gameState = await fetchNewGame();
    console.log("Fetched game state:", gameState);
    renderBoard(gameState);
  } catch (error) {
    console.error("There was a problem with fetching a new game:", error);
  }
}

async function fetchNewGame(): Promise<Game> {
  const response = await fetch("new-game", { method: "POST" });

  if (!response.ok) {
    throw new Error("Could not fetch new chess game");
  }

  return response.json() as Promise<Game>;
}

const gameContainer = document.getElementById("game-container")!;

function renderBoard(gameState: Game) {
  gameContainer.innerHTML = "";
  const boardDiv = createBoardDiv(gameState);
  gameContainer.appendChild(boardDiv);
}

function createBoardDiv(gameState: Game): HTMLDivElement {
  const boardDiv = document.createElement("div");

  gameState.board.forEach((row, rowIndex) => {
    const rowDiv = createRowDiv(row, rowIndex);
    boardDiv.appendChild(rowDiv);
  });

  return boardDiv;
}

function createRowDiv(row: Piece[], rowIndex: number): HTMLDivElement {
  const rowDiv = document.createElement("div");
  rowDiv.classList.add("board-row");

  row.forEach((piece, pieceIndex) => {
    const squareDiv = createSquareDiv(piece, rowIndex, pieceIndex);
    rowDiv.appendChild(squareDiv);
  });

  return rowDiv;
}

function createSquareDiv(
  piece: Piece,
  rowIndex: number,
  pieceIndex: number
): HTMLDivElement {
  const squareDiv = document.createElement("div");
  squareDiv.classList.add("chess-square");

  if ((rowIndex + pieceIndex) % 2 === 0) {
    squareDiv.classList.add("dark-chess-square");
  } else {
    squareDiv.classList.add("light-chess-square");
  }

  squareDiv.addEventListener("click", () => {
    console.log("Square clicked:", { row: rowIndex, col: pieceIndex });
    squareDiv.classList.add("selected-square");
    handleSquareClick({ row: rowIndex, col: pieceIndex });
  });

  const pieceDiv = createPieceDiv(piece);
  squareDiv.appendChild(pieceDiv);

  return squareDiv;
}

function createPieceDiv(piece: Piece): HTMLDivElement {
  const pieceDiv = document.createElement("div");
  pieceDiv.classList.add("chess-piece");

  const pieceType: PieceType = piece.type & 7;
  if (piece.type === PieceType.None) {
    return pieceDiv;
  }
  const pieceColor: Color = piece.type & 24;

  const pieceImg = document.createElement("img");
  pieceImg.src = PieceImages[pieceType]![pieceColor];
  pieceDiv.appendChild(pieceImg);

  return pieceDiv;
}

let selectedSquares: { from?: Square; to?: Square } = {};

function handleSquareClick(square: Square) {
  if (!selectedSquares.from) {
    selectedSquares.from = square;
  } else {
    selectedSquares.to = square;
    const move: Move = {
      fromRow: selectedSquares.from.row,
      fromCol: selectedSquares.from.col,
      toRow: selectedSquares.to.row,
      toCol: selectedSquares.to.col,
    };
    sendMoveRequest(move);
    selectedSquares = {};
  }
}

//// ASYNC FUNCTIONS

async function sendMoveRequest(move: Move) {
  try {
    await makeMove(move);
    const gameState = await currentGameState();
    renderBoard(gameState);
  } catch (error) {
    console.error(error);
  }
}

async function makeMove(move: Move): Promise<string> {
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

async function currentGameState(): Promise<Game> {
  const response = await fetch("/current-state", { method: "GET" });

  if (!response.ok) {
    throw new Error("Could not fetch game state");
  }

  return response.json();
}
