function WelcomeViewModel() {
    var self = this;

    self.isLoad = false;
    self.isOpening = false;
    self.isConnect = ko.observable(false);
    self.btnText = ko.observable('');

    self.title = C.appName + ' ' + C.appVersion
    self.url = C.url

    self.onLoad = function () {
        self.isLoad = true;

        self.onCheck();
    };

    self.onOpen = function () {
        if (!self.isLoad) return;

        self.isOpening = true;
        self.isConnect(true);

        open();

        setTimeout(function () {
            self.isOpening = false;
        }, 3000);
    };

    self.onCheck = function () {
        setTimeout(function () {
            check();
            self.onCheck();
        }, 1500);
    };

    self.onGithub = function () {
        github();
    };

    self.onGitee = function () {
        gitee();
    };

    self.onClose = function () {
        close();
    };

    self.callJs = function (data) {
        if (!self.isLoad) return;

        if (data === 'connect') {
            if (self.isOpening) return;

            self.isConnect(true);
            self.btnText('goman已在浏览器打开');
        } else if (data === 'disconnect') {
            if (self.isOpening) return;

            self.isConnect(false);
            self.btnText('在浏览器中打开goman');
        }
    }
};

var vm = new WelcomeViewModel();
ko.applyBindings(vm);
vm.onLoad();