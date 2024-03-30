/** @type {import('tailwindcss').Config} */
module.exports = {
  // content: [`/templates/**/*_templ.go`],
  content: [`./internal/templates/**/*.templ`],
  theme: {
    extend: {
      boxShadow: {
        "b-outer": "0 2px 3px -2px #8cffd9 , black 0px 0px 10px inset",
        "b-outer2": "black 0px 0px 10px",
      },
      width: {
        cp: "4.5%",
        lg: "55%",
      },
      screens: {
        xs: "480px",
      },
      colors: {
        "black-opac": "#020c12bf",
        "aqua-opac": "#8cffd9bf",
        "aqua-light": "#8cffd9",
        "aqua-light-2": "#d1fff0",
        "aqua-light-3": "#d1fff07e",
        bluer: "#020F19",
        "blue-dk": "#020b11",
      },
    },
  },
  plugins: [],
};