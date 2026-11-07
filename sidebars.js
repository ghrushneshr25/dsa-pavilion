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
          label: 'Problems',
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
      label: 'Linked List',
      link: {
        type: 'doc',
        id: 'linked-list/index',
      },
      items: [
        {
          type: 'category',
          label: 'Concepts',
          items: [
            'linked-list/singly-linked-list',
            'linked-list/doubly-linked-list',
            'linked-list/circular-linked-list',
            'linked-list/xor-linked-list',
            'linked-list/unrolled-linked-list',
            'linked-list/skip-list',
          ],
        },
        {
          type: 'category',
          label: 'Problems',
          items: [
            'linked-list/find-kth-node-from-end-of-linked-list',
          ],
        },
      ],
    },
  ],
};

export default sidebars;
