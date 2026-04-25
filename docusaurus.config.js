import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'DSA Pavilion',
  tagline: 'Data structures & algorithms in Go',
  favicon: 'img/favicon.ico',

  future: {
    v4: true,
  },

  // ✅ CHANGE THIS (critical for GitHub Pages)
  url: 'https://ghrushneshr25.github.io',

  // GitHub Pages project: https://<org>.github.io/<projectName>/
  baseUrl: '/dsa-pavilion/',

  organizationName: 'ghrushneshr25',
  projectName: 'dsa-pavilion',
  trailingSlash: false,
  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.js',
          routeBasePath: '/',
        },

        blog: false,

        theme: {
          customCss: './src/css/custom.css',
        },
      },
    ],
  ],

  themeConfig: {
    image: 'img/docusaurus-social-card.jpg',

    colorMode: {
      respectPrefersColorScheme: true,
    },

    navbar: {
      title: '',
      logo: {
        alt: 'DSA Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'custom-githubSource',
          position: 'right',
          label: 'GitHub',
        },
      ],
    },

    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  },
};

export default config;