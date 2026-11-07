import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'DSA Pavilion',
  tagline: 'Data structures & algorithms in Go',
  favicon: 'img/favicon.svg',

  future: {
    v4: true,
  },

  url: 'https://ghrushneshr25.github.io',
  baseUrl: '/dsa-pavilion/',

  organizationName: 'ghrushneshr25',
  projectName: 'dsa-pavilion',
  trailingSlash: false,
  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  stylesheets: [
    {
      href: 'https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;500;600&family=Instrument+Sans:ital,wght@0,400;0,500;0,600;0,700;1,400&display=swap',
      type: 'text/css',
    },
  ],

  headTags: [
    {
      tagName: 'link',
      attributes: {rel: 'preconnect', href: 'https://fonts.googleapis.com'},
    },
    {
      tagName: 'link',
      attributes: {
        rel: 'preconnect',
        href: 'https://fonts.gstatic.com',
        crossorigin: 'anonymous',
      },
    },
    {
      tagName: 'meta',
      attributes: {name: 'theme-color', content: '#000000'},
    },
  ],

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
    colorMode: {
      defaultMode: 'dark',
      disableSwitch: false,
      respectPrefersColorScheme: false,
    },

    navbar: {
      title: 'pavilion',
      logo: {
        alt: 'DSA',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'custom-search',
          position: 'right',
        },
        {
          type: 'custom-githubSource',
          position: 'right',
          label: 'GitHub',
        },
      ],
    },

    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.vsDark,
    },

    footer: {
      style: 'dark',
      copyright: `<span>Ghrushnesh Rathod</span><span>ctrl.plane · dsa.pavilion</span><a href="https://github.com/ghrushneshr25/dsa" target="_blank" rel="noreferrer">/ghrushneshr25</a>`,
    },
  },
};

export default config;
