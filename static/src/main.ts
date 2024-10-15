// import * as Types from "./types";
// TYPES
const boardsize = 8;

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
  board: Piece[];
  activeColor: Color;
}

interface Move {
  startSquare: number;
  targetSquare: number;
}

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

const gameContainer = document.getElementById("game-container")!;

function renderBoard(gameState: Game) {
  gameContainer.innerHTML = "";
  const boardDiv = createBoardDiv(gameState);
  gameContainer.appendChild(boardDiv);
}

function createBoardDiv(gameState: Game): HTMLDivElement {
  const boardDiv = document.createElement("div");

  for (let rank = 7; rank >= 0; rank--) {
    const rowDiv = document.createElement("div");
    rowDiv.classList.add("board-row");

    for (let file = 0; file < boardsize; file++) {
      const index = rank * boardsize + file;
      const piece = gameState.board[index];

      const squareDiv = createSquareDiv(piece, index);
      squareDiv.dataset.index = index.toString();

      if ((rank + file) % 2 === 0) {
        squareDiv.classList.add("light-chess-square");
      } else {
        squareDiv.classList.add("dark-chess-square");
      }

      rowDiv.appendChild(squareDiv);
    }

    boardDiv.appendChild(rowDiv);
  }

  return boardDiv;
}

function createSquareDiv(piece: Piece, index: number): HTMLDivElement {
  const squareDiv = document.createElement("div");
  squareDiv.classList.add("chess-square");

  squareDiv.textContent = index.toString();

  squareDiv.addEventListener("click", () => {
    handleSquareClick(index);
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

let selectedSquares: number[] = [];

function handleSquareClick(index: number) {
  if (selectedSquares.length === 0) {
    selectedSquares[0] = index;
    getLegalMoves(index);
    return;
  }
  if (selectedSquares[0] === index) {
    selectedSquares = [];
    removeHighlight();
    return;
  }
  selectedSquares[1] = index;
  const move: Move = {
    startSquare: selectedSquares[0],
    targetSquare: selectedSquares[1],
  };
  sendMoveRequest(move);
  selectedSquares = [];
  removeHighlight();
}

function highlightLegalMoves(moves: Move[]) {
  removeHighlight();

  for (let i = 0; i < moves.length; i++) {
    highlightSquare(moves[i].targetSquare);
  }
}

function highlightSquare(index: number) {
  const square = document.querySelector(
    `[data-index='${index}']`
  ) as HTMLElement;
  square.classList.add("highlighted-square");

  const circle = document.createElement("div");
  circle.classList.add("highlight-circle");
  square.appendChild(circle);
}

function removeHighlight() {
  const squares = document.getElementsByClassName("chess-square");
  for (let i = 0; i < squares.length; i++) {
    const square = squares[i] as HTMLElement;
    square.classList.remove("highlighted-square");
    const existingCircle = square.querySelector(".highlight-circle");
    if (existingCircle) {
      existingCircle.remove();
    }
  }
}

//// ASYNC FUNCTIONS

async function startNewGame() {
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

async function getLegalMoves(index: number) {
  try {
    const moves = await fetchLegalMoves(index);
    console.log("Legal moves:", moves);
    highlightLegalMoves(moves);
  } catch (error) {
    console.error("There was a problem with fetching legal moves:", error);
  }
}

async function fetchLegalMoves(index: number): Promise<Move[]> {
  const response = await fetch(`/legal-moves/${index}`, { method: "GET" });

  if (!response.ok) {
    throw new Error("Could not fetch legal moves");
  }

  return response.json();
}
