export default {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go", "./**/*.js" ],
  theme: {
    extend: {
      keyframes: {
        type: {
          "0%": { width: "0%" },
          "100%": { width: "100%" },
        },

        blink: {
          "0%, 100%": { opacity: 1 },
          "50%": { opacity: 0 },
        },

        alertBg: {
          "0%, 100%": { backgroundColor: "theme('colors.gray.200')" },
          "50%": { backgroundColor: "theme('colors.orange.300')" },
        }
      },

      animation: {
        blink: "blink 1s infinite",
        "type-sm": "type 0.5s steps(10, end)",
        "type-lg": "type 1.5s steps(30, end)",
        "alert-bg": "alertBg 1.5s 3",
      }
    }
  },
};
