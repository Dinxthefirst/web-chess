"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const BoardSize = 8;
var Color;
(function (Color) {
    Color["White"] = "White";
    Color["Black"] = "Black";
})(Color || (Color = {}));
function makeMove(move) {
    return __awaiter(this, void 0, void 0, function* () {
        const response = yield fetch("/move", {
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
    });
}
