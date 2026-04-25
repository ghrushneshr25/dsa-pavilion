/**
 * Extends classic theme navbar item registry with custom-githubSource.
 * Spread @theme-original so we never drop built-in types (avoids SSR/client mismatch).
 */
import OriginalComponentTypes from '@theme-original/NavbarItem/ComponentTypes';
import GithubSourceNavbarItem from './GithubSourceNavbarItem';

import type {ComponentTypesObject} from '@theme/NavbarItem/ComponentTypes';

const ComponentTypes: ComponentTypesObject = {
  ...OriginalComponentTypes,
  'custom-githubSource': GithubSourceNavbarItem,
};

export default ComponentTypes;
