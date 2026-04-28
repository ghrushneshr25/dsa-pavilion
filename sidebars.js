// @ts-check
const sidebars = {
  tutorialSidebar: [
    {
      type: 'doc',
      id: 'index',
      label: 'Home',
    },
    {
      type: 'category',
      label: 'Recursion Backtracking',
      link: {
        type: 'doc',
        id: 'recursion-backtracking/index',
      },
      items: [
        {
          type: 'category',
          label: '🧩 Problems',
          items: [
            'recursion-backtracking/towers-of-hanoi',
            'recursion-backtracking/check-if-array-is-sorted-using-recursion',
            'recursion-backtracking/generate-binary-strings',
            'recursion-backtracking/generate-k-ary-strings',
            'recursion-backtracking/largest-connected-region-in-matrix',
          ],
        },
      ],
    },
    {
      type: 'category',
      label: 'Linkedlist',
      link: {
        type: 'doc',
        id: 'linkedlist/index',
      },
      items: [
        {
          type: 'category',
          label: '📘 Concepts',
          items: [
            'linkedlist/singly-linked-list',
            'linkedlist/doubly-linked-list',
            'linkedlist/circular-linked-list',
          ],
        },
      ],
    },
  ],
};

export default sidebars;
