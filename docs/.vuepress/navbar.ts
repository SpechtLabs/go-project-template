import { defineNavbarConfig } from "vuepress-theme-plume";

// The four sections follow Diataxis: Getting Started teaches by doing, Guides
// solve one task each, Understanding explains, and Reference is for lookup.
export const navbar = defineNavbarConfig([
  { text: "Home", link: "/", icon: "mdi:home" },
  { text: "Getting Started", link: "/getting-started/quick-start", icon: "mdi:rocket-launch" },
  { text: "Guides", link: "/guides/install", icon: "mdi:compass" },
  { text: "Understanding", link: "/understanding/overview", icon: "mdi:lightbulb" },
  { text: "Reference", link: "/reference/cli", icon: "mdi:book" },
  {
    text: "More",
    icon: "mdi:dots-horizontal",
    items: [
      {
        text: "Download",
        link: "https://github.com/SpechtLabs/PROJECT_NAME/releases",
        target: "_blank",
        rel: "noopener noreferrer",
        icon: "mdi:download",
      },
      {
        text: "Report an Issue",
        link: "https://github.com/SpechtLabs/PROJECT_NAME/issues/new/choose",
        target: "_blank",
        rel: "noopener noreferrer",
        icon: "mdi:bug-outline",
      },
    ],
  },
]);
