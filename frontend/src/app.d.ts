// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
    interface Piece {
      type: number;
    }
    interface Game {
      board: Piece[];
      ColorToMove: Color;
    }
    interface Move {
      startSquare: number;
      targetSquare: number;
      flag: number;
    }
  }
}

export {};
