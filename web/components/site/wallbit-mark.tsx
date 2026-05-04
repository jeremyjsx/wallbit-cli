type WallbitMarkProps = {
  className?: string;
  title?: string;
};

/**
 * Wallbit brand mark.
 *
 * The source asset paints the OUTER rectangle and uses the "W" as
 * negative space. Rendering it directly on a dark background paints
 * the rectangle in `currentColor` and leaves the W dark — the opposite
 * of what we want.
 *
 * Fix: use an SVG `<mask>` to invert the fill semantics. The mask
 * starts white (= visible) and the original path is painted black
 * (= hidden) over it. We then fill a full-canvas rect with
 * `currentColor` through the mask, so only the negative space of the
 * source path (the W shape) renders.
 *
 * Aspect is 94:112; size with `h-*` and `w-auto`.
 */
export function WallbitMark({ className, title }: WallbitMarkProps) {
  const maskId = "wallbit-mark-cutout";
  const path =
    "M0 560 l0 -560 470 0 470 0 0 135 0 136 -27 -15 c-16 -8 -68 -37 -118 -66 l-90 -52 -60 32 c-33 18 -65 38 -72 44 -9 8 -13 63 -13 185 l-2 173 64 39 c34 21 67 38 72 39 5 0 7 -75 4 -175 -3 -101 -2 -175 3 -175 5 0 27 11 48 24 38 24 38 24 44 108 4 45 7 163 7 261 0 97 3 177 6 177 3 0 35 -16 70 -36 l64 -36 0 161 0 161 -470 0 -470 0 0 -560z m592 433 l118 -67 0 -82 0 -82 -150 -86 c-82 -47 -155 -86 -162 -86 -7 1 -44 18 -82 39 l-68 39 103 55 c57 30 130 68 162 83 31 16 57 32 57 36 0 4 -23 19 -51 33 l-51 26 -186 -107 c-103 -59 -206 -119 -229 -132 l-43 -24 0 80 0 80 227 130 c126 71 230 130 232 131 3 1 58 -29 123 -66z m-282 -458 l149 -86 3 -84 c2 -47 1 -85 -1 -85 -2 0 -66 38 -141 85 -76 47 -145 88 -154 91 -13 5 -16 -2 -16 -48 l0 -55 215 -123 c118 -67 222 -128 230 -134 13 -9 5 -17 -49 -49 -36 -20 -69 -37 -73 -37 -4 0 -109 58 -233 129 l-225 128 -3 136 -3 135 73 41 c40 22 74 41 76 41 1 0 70 -38 152 -85z";

  return (
    <svg
      viewBox="0 0 94 112"
      className={className}
      role={title ? "img" : undefined}
      aria-hidden={title ? undefined : true}
    >
      {title ? <title>{title}</title> : null}
      <defs>
        <mask
          id={maskId}
          maskUnits="userSpaceOnUse"
          maskContentUnits="userSpaceOnUse"
          x="0"
          y="0"
          width="94"
          height="112"
        >
          <rect x="0" y="0" width="94" height="112" fill="white" />
          <g transform="translate(0,112) scale(0.1,-0.1)" fill="black">
            <path d={path} />
          </g>
        </mask>
      </defs>
      <rect
        x="0"
        y="0"
        width="94"
        height="112"
        fill="currentColor"
        mask={`url(#${maskId})`}
      />
    </svg>
  );
}
