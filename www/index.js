// FIXME: unimplemented
import * as simwasm from "simwasm";

const simulation = new simwasm.Simulation();
// --------------- ^-^
// | For all practical purposes, this is a fancy syntax for Rust's:
// | `Simulation::new()`
// ---

const world = simulation.world();
// --------------------- ^---^
// | Parsing already happens inside this automatically-generated
// | function - we don't have to do anything more in here.
// ---

console.log(world);

const viewport = document.getElementById("viewport");
// ------------- ^------^
// | `document` is a global object that allows to access and modify
// | current page (e.g. create or remove stuff from it).
// ---
const viewportWidth = viewport.width; // clientWidth
const viewportHeight = viewport.height; // clientHeight

const viewportScale = window.devicePixelRatio || 1;
// ------------------------------------------ ^^^^
// | Syntax-wise, it's like: .unwrap_or(1)
// |
// | This value determines how much physical pixels there are per
// | each single pixel on a canvas.
// |
// | Non-HiDPI displays usually have a pixel ratio of 1.0, which
// | means that drawing a single pixel on a canvas will lighten-up
// | exactly one physical pixel on the screen.
// |
// | My display has a pixel ratio of 2.0, which means that for each
// | single pixel drawn on a canvas, there will be two physical
// | pixels modified by the browser.
// ---

// The Trick, part 1: we're scaling-up canvas' *buffer*, so that it
// matches the screen's pixel ratio
viewport.width = viewportWidth * viewportScale;
viewport.height = viewportHeight * viewportScale;

// The Trick, part 2: we're scaling-down canvas' *element*, because
// the browser will automatically multiply it by the pixel ratio in
// a moment.
//
// This might seem like a no-op, but the maneuver lies in the fact
// that modifying a canvas' element size doesn't affect the canvas'
// buffer size, which internally *remains* scaled-up:
//
// ----------- < our entire page
// |         |
// |   ---   |
// |   | | < | < our canvas
// |   ---   |   (size: viewport.style.width & viewport.style.height)
// |         |
// -----------
//
// Outside the page, in the web browser's memory:
//
// ----- < our canvas' buffer
// |   | (size: viewport.width & viewport.height)
// |   |
// -----
// TODO:
// v... = ...
// v... = ...

/** @type {CanvasRenderingContext2D} */
const ctx = viewport.getContext("2d");

// Determines color of the upcoming shape.
ctx.fillStyle = "rgb(0, 0, 0);";

for (const animal of simulation.world().animals) {
  ctx.fillRect(animal.x * viewportWidth, animal.y * viewportHeight, 15, 15);
  // ---------- X   Y   W    H
  // | Draws rectangle filled with color determined by `fillStyle`.
  // |
  // | X = position on the X axis (left-to-right)
  // | Y = position on the Y axis (top-to-bottom)
  // | W = width
  // | X = height
  // |
  // | (unit: pixels)
  // ---
}
