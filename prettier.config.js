/*
 * Copyright (c) 2016-present Sven Greb <development@svengreb.de>
 * This source code is licensed under the MIT license found in the license file.
 */

/**
 * Configurations for Prettier.
 * @see https://prettier.io/docs/en/configuration.html
 * @see https://prettier.io/docs/en/options.html
 * @see https://prettier.io/docs/en/options.html#parser
 * @see https://prettier.io/docs/en/plugins.html
 * @see https://github.com/un-ts/prettier/tree/master/packages/sh
 * @see https://github.com/prettier/plugin-xml
 */
module.exports = {
  printWidth: 160,
  overrides: [
    {
      files: ["*.itermcolors", "*.svg"],
      options: {
        parser: "xml",
      },
    },
    {
      files: [".husky/*"],
      options: {
        parser: "sh",
      },
    },
  ],
};
