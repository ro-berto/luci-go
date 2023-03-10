// Copyright 2021 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

////////////////////////////////////////////////////////////////////////////////
// Selector is the dropdown selection menu containing allowlists.
class Selector {
  constructor(element) {
    // Root DOM element.
    this.element = document.getElementById(element);
    // Map of allowlists.
    this.allowlistMap = {};
    // Setup event listener for when selection changes.
    this.element.addEventListener('change', (event) => {
      // TODO(cjacomet): Improve this to not have to iterate over all options.
      // Change which option has the selected attribute.
      this.element.childNodes.forEach((option) => {
        if (!option.dataset.selected && option.text === event.target.value) {
          option.setAttribute('data-selected', 'selected');
        } else {
          option.setAttribute('data-selected', '');
        }
      });
      this.onSelectionChanged(event.target.value);
    });
  }

  populate(allowlists) {
    // Empty the element before populating.
    this.element.innerHTML = '';
    let selected = null;

    let addToSelector = (name) => {
      let option = document.createElement('option');
      option.text = name;
      option.setAttribute('data-allowlist-name', name);
      this.element.appendChild(option);
      if (selected === null) {
        selected = option;
      }
    };

    // All allowlists, populate the selector element.
    allowlists.map((list) => {
      this.allowlistMap[list.name] = list;
      addToSelector(list.name);
    });

    selected.setAttribute('data-selected', 'selected');
    this.onSelectionChanged();
  }

  onSelectionChanged() {
    let selectedOption = this.element.querySelector(
      '[data-selected="selected"]'
    );
    const allowlistSelectedEvent = new CustomEvent('allowlistSelected', {
      bubble: true,
      detail: {
        ip_allowlist: selectedOption.text,
      },
    });

    this.element.dispatchEvent(allowlistSelectedEvent);
  }
}

////////////////////////////////////////////////////////////////////////////////
// AllowlistPane is the panel with information about some selected IP allowlist.
class AllowlistPane {
  constructor() {
    this.descriptionBox = document.getElementById('description-box');
    this.subnetBox = document.getElementById('subnet-box');
  }

  // Fills in the form with details about some IP allowlist.
  populate(ipAllowlist) {
    this.descriptionBox.innerHTML = '';
    this.subnetBox.innerHTML = '';
    this.descriptionBox.value = ipAllowlist.description;
    this.subnetBox.textContent = (ipAllowlist.subnets || []).join('\n');
  }
}

const reloadAllowlists = (selector) => {
  let defer = api.ipAllowlists();
  defer
    .then((response) => {
      selector.populate(response.allowlists);
    })
    .catch((err) => {
      console.log(err);
    });
  return defer;
};

window.onload = () => {
  let selector = new Selector('allowlist-selector');
  let allowlistPane = new AllowlistPane();

  reloadAllowlists(selector);

  selector.element.addEventListener('allowlistSelected', (event) => {
    if (event.detail.ip_allowlist === null) {
      console.log('allowlist is empty');
    } else {
      allowlistPane.populate(selector.allowlistMap[event.detail.ip_allowlist]);
    }
  });
};
