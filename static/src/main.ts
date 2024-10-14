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

function createRowDiv(row: (Piece | null)[], rowIndex: number): HTMLDivElement {
  const rowDiv = document.createElement("div");
  rowDiv.classList.add("board-row");

  row.forEach((piece, pieceIndex) => {
    const squareDiv = createSquareDiv(piece, rowIndex, pieceIndex);
    rowDiv.appendChild(squareDiv);
  });

  return rowDiv;
}

function createSquareDiv(
  piece: Piece | null,
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

function createPieceDiv(piece: Piece | null): HTMLDivElement {
  const pieceDiv = document.createElement("div");
  pieceDiv.classList.add("chess-piece");
  pieceDiv.innerText = piece ? piece.symbol : "";
  return pieceDiv;
}

let selectedSquares: { from?: Square; to?: Square } = {};

function handleSquareClick(square: Square) {
  if (!selectedSquares.from) {
    selectedSquares.from = square;
  } else {
    selectedSquares.to = square;
    sendMoveRequest();
  }
}

async function sendMoveRequest() {
  if (selectedSquares.from && selectedSquares.to) {
    const moveRequest: MoveRequest = {
      fromX: selectedSquares.from.row,
      fromY: selectedSquares.from.col,
      toX: selectedSquares.to.row,
      toY: selectedSquares.to.col,
    };

    try {
      await makeMove(moveRequest);
      const gameState = await currentGameState();
      renderBoard(gameState);
    } catch (error) {
      console.error(error);
    } finally {
      selectedSquares = {};
    }
  }
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

async function currentGameState(): Promise<Game> {
  const response = await fetch("/current-state", { method: "GET" });

  if (!response.ok) {
    throw new Error("Could not fetch game state");
  }

  return response.json();
}
