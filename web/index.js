console.log("from index.js");

import * as simwasm from "./static/js/page.wasm";

const simulation = new simwasm.Simulation();
const world = simulation.World();
// --------------------- ^---^
// | Parsing already happens inside this automatically-generated
// | function - we don't have to do anything more in here.
// ---

console.log(world);

/**
 * Draws a triangle on the canvas.
 *
 * @param {number} x - The X coordinate of the top-left corner of the triangle.
 * @param {number} y - The Y coordinate of the top-left corner of the triangle.
 * @param {number} size - The length of the sides of the triangle.
 * @param {number} rotation - The rotation angle of the triangle.
 * @example ctx.drawTriangle(x*canvas.width, y*canvas.height, 0.01*canvas.width, Math.PI / 4);
 */
CanvasRenderingContext2D.prototype.drawTriangle = function (
    x,
    y,
    size,
    rotation
) {
    /**
     * Triangle is hard to spot when rotated so we extruded one of the vertices.
     * @const {number}
     */
    const extrudeFactor = 1.5;

    this.beginPath();

    this.moveTo(
        x - Math.sin(rotation) * size * extrudeFactor,
        y + Math.cos(rotation) * size * extrudeFactor
    );

    // Instead of + 4.0 / 3.0, we could've also used - 2.0 / 3.0
    // (meaning "60° counterclockwise from the top vertex"):
    this.lineTo(
        x - Math.sin(rotation + (2.0 / 3.0) * Math.PI) * size,
        y + Math.cos(rotation + (2.0 / 3.0) * Math.PI) * size
    );
    this.lineTo(
        x - Math.sin(rotation + (4.0 / 3.0) * Math.PI) * size,
        y + Math.cos(rotation + (4.0 / 3.0) * Math.PI) * size
    );

    this.lineTo(
        x - Math.sin(rotation) * size * extrudeFactor,
        y + Math.cos(rotation) * size * extrudeFactor
    );

    this.stroke(); // FIXME: remove stroke if we are filling?
    this.fillStyle = "rgb(255, 255, 255)"; // white
    this.fill();
};

CanvasRenderingContext2D.prototype.drawCircle = function (x, y, radius) {
    this.beginPath();

    this.arc(x, y, radius, 0, 2.0 * Math.PI);

    this.fillStyle = "rgb(0, 255, 128)"; // green
    this.fill();
};

/** @type {HTMLCanvasElement|null} */
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
viewport.style.width = viewportWidth + "px";
viewport.style.height = viewportHeight + "px";

/** @type {CanvasRenderingContext2D} */
const ctx = viewport.getContext("2d");

// Automatically scales all operations by `viewportScale` - otherwise
// we'd have to `* viewportScale` everything by hand
ctx.scale(viewportScale, viewportScale);

// Determines color of the upcoming shape.
ctx.fillStyle = "rgb(0, 0, 0);";

// FIXME: is this supposed to be in func redraw()?
for (const animal of simulation.World().Animals) {
    ctx.drawTriangle(
        animal.x * viewportWidth,
        animal.y * viewportHeight,
        0.01 * viewportWidth,
        animal.rotation
    );
    // ctx.fillRect(animal.x * viewportWidth, animal.y * viewportHeight, 15, 15);
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

function redraw() {
    ctx.clearRect(0, 0, viewportWidth, viewportHeight);

    simulation.Step();

    const world = simulation.World();

    for (const food of world.Foods) {
        ctx.drawCircle(
            food.X * viewportWidth,
            food.Y * viewportHeight,
            (0.01 / 2.0) * viewportWidth
        );
    }

    for (const animal of world.Animals) {
        ctx.drawTriangle(
            animal.x * viewportWidth,
            animal.y * viewportHeight,
            0.01 * viewportWidth,
            animal.rotation
        );
    }

    // requestAnimationFrame() schedules code only for the next frame.
    //
    // Because we want for our simulation to continue forever, we've
    // gotta keep re-scheduling our function:
    requestAnimationFrame(redraw);
}

redraw();
